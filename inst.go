package goaws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// Host is an entity connected to a network
type Instance struct {
	InstanceId string
	VolumeId   string
	State      ec2.InstanceState
	KeyName    string
	AvailZone  string
	Region     string
	data       interface{}
}

// Instances returns the Instmap
func Instances(region string) map[string]*Instance {
	var reg *Region
	if reg = RegionMap.Get(region); reg == nil {
		return nil
	}
	reg.Instances = FetchInstances(reg.Name)
	return reg.Instances
}

// String returns a single line representing our host
func (i *Instance) String() string {
	return fmt.Sprintf("%s %s %s %s %s",
		i.Region, i.InstanceId, i.VolumeId, i.State, i.KeyName)
}

// GetInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func FetchInstances(region string) (imap map[string]*Instance) {
	e := ec2svc(region)
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	// 4. Parse the response into an instance Map
	imap = imapFromAWS(region, result)
	return imap
}

func imapFromAWS(region string, result *ec2.DescribeInstancesOutput) (imap map[string]*Instance) {

	// Nextoken to read more
	nextToken := result.NextToken
	imap = make(map[string]*Instance)

	resvs := result.Reservations
	for _, resv := range resvs {
		for _, inst := range resv.Instances {
			iid := *inst.InstanceId
			fmt.Printf("inst %+v ", inst)
			os.Exit(1)
			var newinst = &Instance{
				InstanceId: iid,
				State:      *inst.State,
				KeyName:    *inst.KeyName,
				Region:     region,
				data:       inst,
			}

			// Get to the block device mappings
			for _, bdm := range inst.BlockDeviceMappings {
				newinst.VolumeId = *bdm.Ebs.VolumeId
			}
			imap[iid] = newinst
		}
	}
	if nextToken != nil {
		panic("next token != nil ")
	}
	return imap
}

// TerminateInstance will terminate an instance
func TerminateInstances(region string, iids []string) (err error) {

	if iids == nil || len(iids) < 1 {
		for i, inst := range Instances(region) {
			switch inst.State.Name {
			case "terminated":
				// skip this one
			default:
				iids = append(iids, i)
			}
		}
	}

	e := ec2svc(region)
	req := e.TerminateInstancesRequest(&ec2.TerminateInstancesInput{
		InstanceIds: iids,
	})

	result, err := req.Send()
	if err != nil {
		return fmt.Errorf("terminate %v", err)
	}

	for _, tinst := range result.TerminatingInstances {
		iid := *tinst.InstanceId
		fmt.Printf(" %s %s -> %s\n", iid, tinst.PreviousState.Name, tinst.CurrentState.Name)
	}
	return nil
}

package goaws

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// EC2Map
type EC2Instance struct {
	*ec2.Instance
	// TODO updated Time.time // Add the time the item was updated
}

// Instances returns the Instmap
func Instances(region string) map[string]*EC2Instance {
	var reg *Region
	if reg = RegionMap.Get(region); reg == nil {
		return nil
	}
	reg.Instances = FetchInstances(reg.Name)
	return reg.Instances
}

// InstanceID returns InstanceID
func (i *EC2Instance) InstanceID() string {
	return *i.InstanceId
}

// String returns a single line representing our host
func (i *EC2Instance) String() string {
	return fmt.Sprintf("  %s\n    %s %s %s\n", *i.PublicDnsName, *i.InstanceId, i.VolumeId(), i.State.Name)
}

// Region returns the region the instance is in
func (i *EC2Instance) Region() string {
	return *(i.Placement.AvailabilityZone)
}

// Id of this type
func (i *EC2Instance) Id() string {
	return *(i.InstanceId)
}

// PublicDnsName is self descriptive
func (i *EC2Instance) PublicDNSName() string {
	return *(i.PublicDnsName)
}

// Volumes returns all VolumeId's used by this instance
func (i *EC2Instance) Volumes() (volids []string) {
	for _, bdm := range i.BlockDeviceMappings {
		volids = append(volids, *bdm.Ebs.VolumeId)
	}
	return volids
}

// VolumeId returns the volume id of the first volume
func (i *EC2Instance) VolumeId() (volid string) {
	vols := i.Volumes()
	if vols != nil {
		//volid = vols[0].VolumeId
		volid = vols[0]
	}
	return volid
}

// GetInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func FetchInstances(region string) (imap map[string]*EC2Instance) {
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

// Create an insdex of EC2Instances indexed by InstanceId
func imapFromAWS(region string, result *ec2.DescribeInstancesOutput) (imap map[string]*EC2Instance) {
	// Nextoken to read more
	nextToken := result.NextToken
	imap = make(map[string]*EC2Instance)
	resvs := result.Reservations
	for _, resv := range resvs {
		for _, inst := range resv.Instances {
			ecinst := &EC2Instance{
				Instance: &inst,
			}
			imap[*(inst.InstanceId)] = ecinst
		}
	}
	if nextToken != nil {
		panic("next token != nil ")
	}
	return imap
}

// TerminateInstances will terminate an instance
func TerminateInstances(region string, iids []string) (err error) {

	if iids == nil || len(iids) < 1 {
		for _, inst := range Instances(region) {
			log.Fatalf(" %+v ", inst)
			/*
				switch *inst.State {
				case "terminated":
					// skip this one
				default:
					iids = append(iids, i)
				}
			*/
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

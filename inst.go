package goaws

import (
	"fmt"
	"strings"

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
}

// String returns a single line representing our host
func (i *Instance) String() string {
	return fmt.Sprintf("%s %s %s %s %s",
		i.Region, i.InstanceId, i.VolumeId, i.State, i.KeyName)
}

// FetchInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func (cl *AWSCloud) GetInstances() (imap Instmap) {

	if cl.imap != nil {
		return cl.imap
	}

	// 1. Look for a cached version of the object, return if found
	idxname := cl.region + "-inst"
	err := cache.FetchObject(idxname, &imap)
	if err == nil && imap != nil {
		log.Debugf("  found cached version of %s .. ", idxname)
		return imap
	}

	// 3. Prepare and send the AWS request and wait for a response
	// 4. Prepare and send a describe request
	log.Debugf("  fetch instance data from AWS %s ", cl.region)
	e := cl.Client()
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	// 4. Parse the response into an instance Map
	if cl.imap = imapFromAWS(result, cl.region); imap == nil {
		log.Infoln("  failed to get imap from AWS ")
		return nil
	}

	// 5. Store the object for later usage
	go func() {
		obj, err := cache.StoreObject(idxname, imap)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
		}
		log.Debugf("  object cached at path %s ", obj.Path)
	}()

	// 6. Give the caller what they want and return
	return imap
}

// Instance will return the instance with the given IID
func (cl *AWSCloud) Instance(iid string) *Instance {
	if inst, e := cl.imap[iid]; e {
		return inst
	}
	return nil
}

// Create an InstanceMap from the AWS EC2 response
func imapFromAWS(result *ec2.DescribeInstancesOutput, region string) (imap Instmap) {

	// Nextoken to read more
	nextToken := result.NextToken
	imap = make(Instmap)
	resvs := result.Reservations
	for _, resv := range resvs {
		for _, inst := range resv.Instances {

			iid := *inst.InstanceId
			var newinst = &Instance{
				InstanceId: iid,
				State:      *inst.State,
				KeyName:    *inst.KeyName,
				Region:     region,
			}
			for _, bdm := range inst.BlockDeviceMappings {
				newinst.VolumeId = *bdm.Ebs.VolumeId
			}
			imap[iid] = newinst
			I2R[iid] = region
		}
	}

	if nextToken != nil {
		panic("next token != nil ")
	}
	fmt.Println("  returning from imap from aws")
	return imap
}

// TerminateInstance will terminate an instance
func (cl *AWSCloud) TerminateInstances(iids []string) error {
	var e *ec2.EC2

	for i, inst := range cl.imap {
		if strings.Compare(string(inst.State.Name), "terminated") != 0 {
			fmt.Printf("  terminate %s -> %s\n", inst.InstanceId, inst.State.Name)
			iids = append(iids, i)
			if len(iids) > 4 {
				break
			}
		}
	}

	// Create the TerminateInstanceRequest
	req := e.TerminateInstancesRequest(&ec2.TerminateInstancesInput{
		InstanceIds: iids,
	})

	// Send the TIR
	if result, err := req.Send(); err != nil {
		return fmt.Errorf("  failed request instances %v ", err)
	} else {
		log.Info("terminate requests %+v", result)
	}
	return nil
}

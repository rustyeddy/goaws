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
func GetInstances(region string) (imap Instmap) {
	var e *ec2.EC2

	log.Debugln("~~> GetInstances for region ", region)
	defer log.Debugln("  <~~ return GetInstances ", region)

	// 1. Look for a cached version of the object, return if found
	idxname := region + "-inst"
	err := cache.FetchObject(idxname, &imap)
	if err == nil && imap != nil {
		log.Debugf("  found cached version of %s .. ", idxname)
		return imap
	}

	// 2. Get the ec2 client for the specified region
	if e = getEC2(region); e == nil {
		log.Errorf("  failed to get an EC2 client for ", region)
		return nil
	}

	// 3. Prepare and send the AWS request and wait for a response
	log.Debugf("  fetch instance data from AWS %s ", region)

	// 4. Prepare and send a describe request
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	// 4. Parse the response into an instance Map
	if imap = imapFromAWS(result, region); imap == nil {
		log.Errorf("  failed to get imap from AWS %v", err)
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

// GetInstances will get the instances for this cloud
func (cl *AWSCloud) GetInstances(iids []string) Instmap {
	if cl.imap == nil {
		cl.imap = GetInstances(cl.region)
	}
	return cl.imap
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

// GetInstance from InstanceId
func GetInstance(iid string) *Instance {
	if inst, e := AllInstances[iid]; e {
		return inst
	}
	return nil
}

// TerminateInstance will terminate an instance
func TerminateInstances(region string, iids []string) error {
	var (
		e *ec2.EC2
	)
	log.Debugln("~~> TerminateInstance instance id %v ", iids)
	defer log.Debugln("  <~~ return TerminateInstance %v ", iids)

	// 1. Get the ec2 client for the specified region
	if e = getEC2(region); e == nil {
		return fmt.Errorf("expected ec2 cli for (%s) got () ", region)
	}

	// 2. Prepare and send the AWS request and wait for a response
	log.Debugf("  fetch instance data from AWS %s ", region)

	var dryrun *bool
	dr := false
	dryrun = &dr

	if cl := GetCloud(region); cl != nil {
		cl.imap = GetInstances(region)
		for i, inst := range cl.imap {
			if strings.Compare(string(inst.State.Name), "terminated") != 0 {
				fmt.Printf("terminating %s -> %s\n", inst.InstanceId, inst.State.Name)
				iids = append(iids, i)
			}
		}
	}
	req := e.TerminateInstancesRequest(&ec2.TerminateInstancesInput{
		DryRun:      dryrun,
		InstanceIds: iids,
	})

	// Send the terminate instance request
	if result, err := req.Send(); err != nil {
		return fmt.Errorf("  failed request instances %v ", err)
	} else {
		log.Fatalf(" %+v ", result)
	}
	return nil
}

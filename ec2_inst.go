package goaws

import (
	"fmt"

	log "github.com/rustyeddy/logrus"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Some fundamental types
type Instmap map[string]*Instance

// Host is an entity connected to a network
type Instance struct {
	InstanceId string
	VolumeId   string
	State      ec2.InstanceState
	KeyName    string
	AvaillZone string
	Region     string
}

// String returns a single line representing our host
func (i *Instance) String() string {
	return fmt.Sprintf("%s %s %s %s", i.InstanceId, i.VolumeId, i.State, i.KeyName)
}

// FetchInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func GetInstances(region string) (imap Instmap) {
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
	e := getEC2(region)
	if e == nil {
		log.Errorf("  failed to get an EC2 client for ", region)
		return nil
	}

	// 3. Prepare and send the AWS request and wait for a response
	log.Debugf("  fetch instance data from AWS %s ", region)
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err := req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	// 4. Parse the response into an instance Map
	if imap = imapFromAWS(result); imap == nil {
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

// Create an InstanceMap from the AWS EC2 response
func imapFromAWS(result *ec2.DescribeInstancesOutput) (imap Instmap) {

	// Nextoken to read more
	nextToken := result.NextToken
	resvs := result.Reservations
	for _, resv := range resvs {
		for _, inst := range resv.Instances {
			var newinst = &Instance{
				InstanceId: *inst.InstanceId,
				State:      *inst.State,
				KeyName:    *inst.KeyName,
			}
			for _, bdm := range inst.BlockDeviceMappings {
				newinst.VolumeId = *bdm.Ebs.VolumeId
			}
			imap[newinst.InstanceId] = newinst
		}
	}
	if nextToken != nil {
		panic("next token != nil ")
	}
	return imap
}

// DeleteInstance
func DeleteInstance(instId string) error {

	/*

		var svc *ec2.EC2
		if svc = getEC2(region); svc == nil {
			log.Errorf("  failed to get aws client for %s ", region)
			return nil
		}

		log.Debugf("  sending request to delete instance region ", region)
		req := svc.DeleteInstanceRequest(&ec2.DeleteInstanceInput{VolumeId: aws.String(instId)})
		result, err := req.Send()
		if err != nil {
			return fmt.Errorf("  # failed response to request %v", err)
		}
		log.Debugf("  got result %v from region %s ", result, region)
		log.Fatalf("  result %+v", result)
	*/
	return nil
}

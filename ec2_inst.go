package goaws

import (
	"fmt"

	log "github.com/rustyeddy/logrus"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// FetchInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func FetchInstances(region string) (result *ec2.DescribeInstancesOutput) {
	log.Println("~~> GetInstances for region ", region)
	defer log.Println("  <~~ return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	e := getEC2(region)
	if e == nil {
		log.Errorf("  failed to get an EC2 client for ", region)
		return nil
	}

	// Look for a cached version of the object
	idxname := region + "-instances"
	err := cache.FetchObject(idxname, result)
	if err == nil && result != nil {
		log.Debugf("  found cached version of %s .. ", idxname)
		return result
	}

	log.Debugf("  fetch instance data from AWS %s ", region)
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err = req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	// Store the object for later cache usage
	go func() {
		obj, err := cache.StoreObject(idxname, result)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
			return
		}
		log.Debugf("  object cached at path %s ", obj.Path)
	}()

	fmt.Printf(" %+v ", result)

	// Parse the json and see what we got
	// vms := ParseInstances()

	return result
}

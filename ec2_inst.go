package goaws

import (
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

	var err error
	var idxname string

	caching := false
	if caching {
		idxname = region + "-instances"
		log.Debugf("  instances idx %s", idxname)

		// Check our local cache first
		err := cache.FetchObject(idxname, result)
		if err == nil && result != nil {
			log.Debugf("  found cached version of %s .. ", idxname)
			return result
		}
	}

	log.Debugf("  fetch instance data from AWS %s ", region)
	req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
	result, err = req.Send()
	if err != nil {
		log.Errorf("  failed request instances %v ", err)
		return nil
	}

	if caching {
		log.Debug("  AWS fetch successful, store object in cache .. ")
		obj, err := cache.StoreObject(idxname, result)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
			return nil
		}
		log.Debugf("  object cached at path %s ", obj.Path)
	}
	return result
}

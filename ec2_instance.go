package goaws

import (
	log "github.com/rustyeddy/logrus"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// FetchInstances will retrieve instances from AWS, it will also store
// the results in the Object cache as a JSON file.
func FetchInstances(region string) (result *ec2.DescribeInstancesOutput) {
	log.Println("   GetInstances for region ", region)
	defer log.Println("  return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	if e := GetEC2(region); e != nil {

		idxname := region + "-instances"

		// Check our local cache first
		err := cache.FetchObject(idxname, result)
		if err != nil {
			log.Debugf("  could not find idx: %s ", idxname)
		} else if result == nil {
			log.Errorf("  failed to get index %s ", idxname)
		} else {
			log.Debugf("  returning cached version of results ")
			return result
		}

		log.Debugf("  fetch instance data from AWS  ", region)
		req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
		result, err = req.Send()
		if err != nil {
			log.Errorf("  failed request instances %+v ", err)
			return nil
		}

		log.Debug("  got results from AWS fetch, store object in cache ")
		obj, err := cache.StoreObject(idxname, result)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
			return nil
		}
		log.Debug("  cached object in path %s ", obj.Path)
	} else {
		log.Fatalf("  failed to get EC2 client for region %s ", region)
	}
	return result
}

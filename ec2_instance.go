package goaws

import (
	log "github.com/rustyeddy/logrus"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// FetchInstances will retrieve instances from AWS
func FetchInstances(region string) *ec2.DescribeInstancesOutput {
	log.Println("   GetInstances for region ", region)
	defer log.Println("  return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	if e := GetEC2(region); e != nil {
		req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
		result, err := req.Send()
		if err != nil {
			log.Errorf("  failed request instances %+v ", err)
			return nil
		}

		idxname := region + "-instances"
		obj, err := cache.StoreObject(idxname, result)
		if err != nil {
			log.Errorf("  failed to store object %s -> err ", idxname, err)
			return nil
		}
		log.Debug("  got an object %+v ", obj)
		return result
	} else {
		log.Fatalf("  failed to get EC2 client for region %s ", region)
	}
	return nil
}

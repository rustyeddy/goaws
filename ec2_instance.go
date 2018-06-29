package goaws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// FetchInstances will retrieve instances from AWS
func FetchInstances(region string) *ec2.DescribeInstancesOutput {
	log.Println("   GetInstances for region ", region)
	defer log.Println("  return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	if e := GetEC2(region); e != nil {
		req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
		if result, err := req.Send(); err == nil {
			log.Fatalf("result => %+v ", result)
			// Index the Instances we've recieved
			// log.Debugln("  index Instances ... ")
			// inv.indexInstances(result.Reservations)

			// log.Debugln("  save Instances ... ")
			// inv.saveInstances(result.Reservations)
			return result
		}
	} else {
		log.Fatalf("  failed to get EC2 client for region %s ", region)
	}
	return nil
}

package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchInstances will retrieve instances from AWS
func FetchInstances(region string) {
	log.Println("   GetInstances for region ", region)
	defer log.Println("  return GetInstances ", region)

	// Fetch the inventory for this region from AWS
	if e := GetEC2(region); e != nil {
		req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
		if result, err := req.Send(); err == nil {

			log.Printf("result => %+v", result)

			// Index the Instances we've recieved
			log.Debugln("  index Instances ... ")
			//inv.indexInstances(result.Reservations)

			log.Debugln("  save Instances ... ")
			//inv.saveInstances(result.Reservations)
		}
	}
}

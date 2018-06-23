package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchInstances will retrieve instances from AWS
func (inv *Inventory) FetchInstances() {
	log.Println("GetInstances for region ", inv.Region)
	defer log.Println("  return GetInstances ", inv.Region)

	// Fetch the inventory for this region from AWS
	if e := inv.GetEC2(); e != nil {
		req := e.DescribeInstancesRequest(&ec2.DescribeInstancesInput{})
		if result, err := req.Send(); err == nil {
			// Index the Instances we've recieved
			inv.indexInstances(result.Reservations)
			inv.saveInstances(result.Reservations)
		}
	}
}

// saveInstances will save Instances from AWS to json file.
func (inv *Inventory) saveInstances(res []ec2.RunInstancesOutput) {
	if _, err := S.StoreObject("instances", res); err != nil {
		log.Errorf("StoreObject failed ", err)
	}
}

// DeleteInstance
func (inv *Inventory) deleteInstance(iid string) {
	if err := S.RemoveObject(iid); err != nil {
		log.Errorf("failed to remove object %s -> %v", iid, err)
	}
	log.Debugf("removed object %s ", iid)
}

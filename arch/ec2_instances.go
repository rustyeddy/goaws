package goaws

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// saveInstances will save Instances from AWS to json file.
func (inv *Inventory) saveInstances(res []ec2.RunInstancesOutput) {
	if _, err := cache.StoreObject("instances", res); err != nil {
		log.Errorf("StoreObject failed %v", err)
	}
}

// DeleteInstance
func (inv *Inventory) deleteInstance(iid string) {
	if err := cache.RemoveObject(iid); err != nil {
		log.Errorf("failed to remove object %s -> %v", iid, err)
	}
	log.Debugf("removed object %s ", iid)
}

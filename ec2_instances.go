package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchInstances will retrieve instances from aws
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

// SaveInstances will save Instances from AWS to json file.
func (inv *Inventory) saveInstances(res []ec2.RunInstancesOutput) {
	// Cache the results in a local file
	jbytes, err := json.Marshal(res)
	if err != nil {
		log.Errorf("describe instances request %v", err)
		return
	}
	// Cache a local copy of the string
	fname := "run/inst-" + inv.Region + ".json"
	if err = ioutil.WriteFile(fname, jbytes, 0644); err != nil {
		log.Error("failed to save ", fname, err)
	}
}

// DeleteInstance
func (inv *Inventory) deleteInstance(iid string) {
	panic("TODO")
}

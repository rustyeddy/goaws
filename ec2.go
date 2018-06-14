package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

/*
   These functions handles all communication with aws. Xlate files between JSON & Go.
   This file also handles caching aws data in local files system.

   Caching is simple.  Get requested data (regions, volumes or instances),
   the get function checks local filesystem, then heads to provider (aws)
   for real data.

   Once data has been read, it is written to local filesystem as JSON for
   future reference.

   To Renew data simply delete cached JSON files.
*/

// GetEC2 returns an ec2 service ready for use
func GetEC2(name string) *ec2.EC2 {

	cfg, err := external.LoadDefaultAWSConfig()
	if err == nil {
		log.Error("failed to get an EC2 client ", name, err)
		return nil
	}
	cfg.Region = name
	ec2Svc := ec2.New(cfg)
	if ec2Svc == nil {
		log.Error("failed to get an EC2 client ", name, err)
	}
	return ec2Svc
}

// GetEC2 Inventory specific version of EC2 client
func (inv *Inventory) GetEC2() *ec2.EC2 {
	if inv.EC2 == nil {
		inv.EC2 = GetEC2(inv.Region)
	}
	return inv.EC2
}

// FetchInventories gather instance and volume data from all AWS regions.
func FetchInventories() string {
	regions := Regions()
	if regions == nil {
		log.Fatalf("regions should not be nil %+v", regions)
	}
	for _, region := range regions {
		inv := inventories.Get(region)
		if inv != nil {
			inv.FetchInstances()
			inv.FetchVolumes()
		}
	}
	return "Fetch inventories complete"
}

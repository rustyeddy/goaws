package goaws

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

/* These functions handle all communication with aws. Translate files
   between JSON & Go.  This file also handles caching aws data in
   local files system.

   Caching is simple.  Get requested data (regions, volumes or
   instances), the get function checks local filesystem, then heads to
   provider (aws) for real data.

   Once data has been read, it is written to local filesystem as JSON
   for future reference.

   To Renew data simply delete cached JSON files.
*/

// GetEC2 returns an ec2 service ready for use
func GetEC2(name string) (ec2Svc *ec2.EC2) {
	name = "us-west-1"
	log.Debugln("Get EC2 for region ", region)
	defer log.Debugln(" leaving EC2 %+v ", ec2Svc)

	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("  Failed to Load Default AWS Config %q -> %v ", name, err)
		return nil
	}

	log.Debugf("  loaded Default config, create ec2 client ")

	//cfg.Region = endpoints.UsWest2RegionID
	ec2Svc = ec2.New(cfg)
	if ec2Svc == nil {
		log.Fatalln("failed to get an EC2 client ", name, err)
	}
	return ec2Svc
}

// GetEC2 Inventory specific version of EC2 client
/*
func (inv *Inventory) GetEC2() *ec2.EC2 {
	if inv.EC2 == nil {
		inv.EC2 = GetEC2(inv.Region)
	}
	return inv.EC2
}
*/
// FetchInventories gather instance and volume data from all AWS regions.
/*
func FetchInventories() error {

	log.Debugf("~~> FetchInventories ")
	defer log.Debugf("<~~ return FetchInventories .. ")

	log.Debug("  -- Get regions .. ")
	regions := Regions()
	if regions == nil {
		log.Fatalf("regions should not be nil %+v", regions)
	}

	log.Debugf("  -- walk (%d) regions .. ", len(regions))
	for _, region := range regions {
		inv, e := inventories[region]
		if !e {
			if inv = NewInventory(region); inv == nil {
				log.Fatalf("  ## NewInventory failed .. nil ")
				continue
			}
		}
		if inv == nil {
			return fmt.Errorf("failed inventory for region %s", region)
		}
		inv.FetchInventory()
		return fmt.Errorf("  failed to recieve inventory for %+v ", inv)
	}
	return nil
}

// FetchInventory is specific to a region
func (inv *Inventory) FetchInventory() *Inventory {
	inv.FetchInstances()
	inv.FetchVolumes()
	return inv
}
*/

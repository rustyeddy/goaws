package cmd

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	regions   []string
	instances map[string]*ec2.DescribeInstancesOutput
	volumes   map[string]*ec2.CreateVolumeOutput

	ec2Cmd = cobra.Command{
		Use:   "instances",
		Short: "list EC2 instances",
		Run:   DoEC2,
	}
)

func init() {

	RootCmd.AddCommand(&ec2Cmd)
	instances = make(map[string]*ec2.DescribeInstancesOutput)
	volumes = make(map[string]*ec2.CreateVolumeOutput)
}

// DoEC2 executes the EC2
func DoEC2(cmd *cobra.Command, args []string) {

	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  expected list of regions got ()")
	}

	for _, region := range regions {
		results := goaws.FetchInstances(region)
		if results == nil {
			log.Errorf("  no instances for region %s", region)
			continue
		}

		// Stick this thing on the Instances map
		log.Debugln("  stick the result on the instances map for ", region)
		instances[region] = results

		// Print some instances information
		nextToken := results.NextToken
		resvs := results.Reservations
		for _, resv := range resvs {
			for j, inst := range resv.Instances {
				log.Debugf("  %d instance %s ", j, *inst.InstanceId)
			}
		}
		if nextToken != nil {
			log.Debugf("  nextToken %s -> cnt instances %d ",
				*nextToken, len(instances))
		} else {
			log.Debugf("  instances %d ", len(instances))
		}
	}

	// Lets store the instances in a JSON file
	cache := goaws.Cache()
	if cache == nil {
		log.Error("  failed to get the cache, moving on .. ")
		return
	}
	o, err := cache.StoreObject("instances", instances)
	if err != nil {
		log.Errorf("  failed to store instances %v ", err)
	}
	log.Infof("  cached stored %s ", o.Path)
}

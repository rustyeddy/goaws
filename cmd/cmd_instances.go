package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	regions     []string
	instances   map[string]*ec2.DescribeInstancesOutput
	instanceCmd = cobra.Command{
		Use:   "instances",
		Short: "list instances",
		Run:   doInstances,
	}
)

func init() {
	RootCmd.AddCommand(&instanceCmd)
	instances = make(map[string]*ec2.DescribeInstancesOutput)
}

// DoEC2 executes the EC2
func doInstances(cmd *cobra.Command, args []string) {

	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  expected list of regions got ()")
	}

	for _, region := range regions {

		fmt.Println("Fetching instances for region ", region)
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
				fmt.Printf("  %d instance %s ", j, *inst.InstanceId)
			}
		}

		if nextToken != nil {
			log.Debugf("  nextToken %s -> cnt instances %d ",
				*nextToken, len(instances))
		} else {
			log.Debugf("  instances %d ", len(instances))
		}
	}
	log.Infof("  cached stored %s ", o.Path)
}

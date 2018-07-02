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
		fmt.Printf("\nFetching instances for region %s \n", region)
		results := goaws.FetchInstances(region)
		if results == nil {
			fmt.Println("Failed to fetch instances ... ")
			continue
		}
		log.Fatal("foo")
		instances[region] = results
		nextToken := results.NextToken
		resvs := results.Reservations

		for _, resv := range resvs {
			for _, inst := range resv.Instances {
				fmt.Printf("%s %s\n", *inst.InstanceId, *inst.KeyName)
			}
		}

		if nextToken != nil {
			fmt.Printf("\tnextToken %s -> cnt instances %d\n ",
				*nextToken, len(instances))
		} else {
			fmt.Printf("\tinstances %d\n", len(instances))
		}
	}
	fmt.Println("done...")
}

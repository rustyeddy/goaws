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
	cmdInstance = cobra.Command{
		Use:   "instances",
		Short: "list instances",
		Run:   doInstances,
	}

	cmdDeleteInstance = cobra.Command{
		Use:   "rminst",
		Short: "Remove an instance and associated volume(s) if any",
		Run:   doDeleteInstance,
	}
)

func init() {
	instances = make(map[string]*ec2.DescribeInstancesOutput)
	RootCmd.AddCommand(&cmdInstance)
	RootCmd.AddCommand(&cmdDeleteInstance)
}

// DoEC2 executes the EC2
func doInstances(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  expected list of regions got ()")
	}

	for _, region := range regions {
		// fmt.Printf("\nFetching instances for region %s \n", region)
		instances := goaws.GetInstances(region)
		if instances == nil {
			// fmt.Println("Failed to fetch instances ... ")
			continue
		}
		for _, inst := range instances {
			fmt.Println(inst)
		}
	}
}

func doDeleteInstance(cmd *cobra.Command, args []string) {
	// Find an instance and try to delete it..
	regions := goaws.Regions()
	for _, region := range regions {
		fmt.Printf("Try to delete something from %s\n ", region)
		vols := goaws.GetVolumes(region)
		if vols == nil {
			log.Fatalf("Expected Volumes got NOTHING")
		}

		for _, vol := range vols {
			if err := goaws.DeleteVolume(region, vol.VolumeId); err != nil {
				log.Fatalf("  failed to send DeleteVolume request %v", err)
			}
			log.Fatalf("VOL:  %+v ", vol)
		}
	}
}

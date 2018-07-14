package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	// InstCmd goa instances
	InstanceCmd = cobra.Command{
		Use:   "instance",
		Short: "Manage AWS Instances",
		Run:   cmdInstances,
	}

	DescribeInstanceCmd = cobra.Command{
		Use:   "describe",
		Short: "Describe the instance by instance-Id",
		Run:   cmdDescribeInstance,
	}

	// InstCmd goa instances
	TerminateInstancesCmd = cobra.Command{
		Use:   "terminate",
		Short: "Manage AWS Instances",
		Run:   cmdTerminateInstances,
	}
)

func init() {
	// go instance terminate iid iid2 iid3 ...
	InstanceCmd.AddCommand(&TerminateInstancesCmd)
}

// List the instances
func cmdInstances(cmd *cobra.Command, args []string) {
	var regions []string

	if Config.Region != "" {
		regions = append(regions, Config.Region)
	} else {
		if regions = goaws.Regions(); regions == nil {
			log.Fatal("  expected (regions) got ()")
		}
	}
	log.Debugln("regions %+v", regions)

	// Walk the regions
	for _, region := range regions {
		fmt.Printf("Instances for region %s ... \n ", region)
		cl := goaws.GetCloud(region)
		for _, inst := range cl.Instances() {
			fmt.Printf("    %s %s %s \n", inst.InstanceId, inst.VolumeId, inst.State.Name)
		}
	}
}

// Describe Instances
func cmdDescribeInstance(cmd *cobra.Command, args []string) {
	inst := goaws.GetInstance(args[0])
	fmt.Printf("instnace %+v", inst)
}

// Terminate Instances
func cmdTerminateInstances(cmd *cobra.Command, args []string) {
	if err := goaws.TerminateInstances(Config.Region, args); err != nil {
		log.Errorf("  problems with Terminate Instances %v", err)
		return
	}
}

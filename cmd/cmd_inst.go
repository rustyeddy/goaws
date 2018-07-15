package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	// InstCmd goa instances
	instanceCmd = cobra.Command{
		Use:     "instance",
		Aliases: []string{"inst"},
		Short:   "Manage AWS Instances",
		Run:     cmdInstances,
	}

	describeInstanceCmd = cobra.Command{
		Use:     "describe",
		Aliases: []string{"desc", "info"},
		Short:   "Describe the instance by instance-Id",
		Run:     cmdDescribeInstance,
	}

	terminateInstancesCmd = cobra.Command{
		Use:     "terminate",
		Aliases: []string{"kill"},
		Short:   "Manage AWS Instances",
		Run:     cmdTerminateInstances,
	}
)

func init() {
	// go instance terminate iid iid2 iid3 ...
	instanceCmd.AddCommand(&terminateInstancesCmd)
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
	cl := goaws.GetCloud(Config.Region)
	inst := cl.Instance(args[0])
	fmt.Printf("instnace %+v", inst)
}

// Terminate Instances
func cmdTerminateInstances(cmd *cobra.Command, args []string) {

	// Get our cloud
	cl := goaws.GetCloud(Config.Region)
	if cl != nil {
		log.Errorf("expected cloud (%s) got () ", Config.Region)
	}

	// Request to actually terminate
	if err := cl.TerminateInstances(args); err != nil {
		log.Errorf("  problems with Terminate Instances %v", err)
		return
	}
}

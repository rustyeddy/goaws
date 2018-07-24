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
		Aliases: []string{"inst", "instances"},
		Short:   "Manage AWS Instances",
		Run:     cmdInstances,
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
	var regionNames []string
	if regionNames = GetRegions(cmd, args); regionNames == nil {
		log.Fatal("expected (regions) got () ")
	}
	for _, region := range regionNames {
		fmt.Printf("Instances for region %s ... \n ", region)
		for _, inst := range goaws.Instances(region) {
			fmt.Printf("  %s %s %s \n", *inst.InstanceId, inst.VolumeId(), inst.State.Name)
		}
	}
	fmt.Println("finished..")
}

// Describe Instances
func cmdDescribeInstance(cmd *cobra.Command, args []string) {

	insts := goaws.Instances(Config.Region)
	inst, e := insts[args[0]]
	if e {
		fmt.Printf("instnace %+v", inst)
	} else {
		fmt.Printf("instnace %s not found", args[0])
	}
}

// Terminate Instances
func cmdTerminateInstances(cmd *cobra.Command, args []string) {

	var (
		regions []string
	)

	if len(args) > 0 {
		regions = args
	} else {
		if regions = goaws.Regions(); regions == nil {
			log.Fatalf("  failed to get regions")
		}
	}

	for _, region := range regions {
		instances := goaws.Instances(region)
		var iids []string
		for n := range instances {
			iids = append(iids, n)
		}
		icount := len(iids)
		if icount > 0 {
			fmt.Printf("Terminate %d instances in region %s\n", icount, region)
			fmt.Printf("  terminating %s, n")
			if err := goaws.TerminateInstances(region, iids); err != nil {
				log.Errorf("  problems with Terminate Instances %v", err)
				return
			}
		} else {
			fmt.Println("  no instances to be terminated ...")
		}

	}
}

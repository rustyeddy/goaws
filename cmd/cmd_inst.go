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
	var regs []string
	var err error

	if Config.Region != "" {
		regs = append(regs, Config.Region)
	} else {
		if regs, err = goaws.Regions(); regs == nil {
			log.Fatal("  expected (regions) got (%v)", err)
		}
	}

	output := ""
	for _, region := range regs {
		fmt.Sprintf(output, "Instances for region %s ... \n ", region)
		for _, inst := range goaws.Instances(region) {
			fmt.Sprintf(output, "  %s %s %s \n", inst.InstanceId, inst.VolumeId, inst.State.Name)
		}
	}
	fmt.Println(output)
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

	// Request to actually terminate
	if err := goaws.TerminateInstances(Config.Region, args); err != nil {
		log.Errorf("  problems with Terminate Instances %v", err)
		return
	}
}

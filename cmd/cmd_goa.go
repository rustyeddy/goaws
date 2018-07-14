/*
Goa commands aws management utilities.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	cache *goaws.Cache // store.Store

	// GoaCmd is the root command
	GoaCmd = cobra.Command{
		Use:   "goa",
		Short: "Manage AWS Instances and Volumes",
		Long:  `Gather AWS Volume and Image information, do stuff with it`,
		Run:   cmdGoa,
	}

	// RegionsCmd list regions
	RegionCmd = cobra.Command{
		Use:   "region",
		Short: "List AWS Regions",
		Run:   cmdRegions,
	}

	// InstCmd goa instances
	InstanceCmd = cobra.Command{
		Use:   "instance",
		Short: "Manage AWS Instances",
		Run:   cmdInstances,
	}

	// InstCmd goa instances
	TerminateInstancesCmd = cobra.Command{
		Use:   "terminate",
		Short: "Manage AWS Instances",
		Run:   cmdTerminateInstances,
	}

	// VolCmd list Volumes from all regions (or the -region flag)
	VolumeCmd = cobra.Command{
		Use:   "volume",
		Short: "Manage AWS Volumes",
		Run:   cmdVolumes,
	}

	// VolDeleteCmd will delete the given volume
	VolumeDeleteCmd = cobra.Command{
		Use:   "delete",
		Short: "delete the volume by vol-Id",
		Run:   cmdDeleteVolume,
	}

	// SnapCmd list snap shots
	SnapshotCmd = cobra.Command{
		Use:   "snap",
		Short: "Manage Host Snaphosts",
		Run:   cmdSnapshots,
	}
)

// Get the AWS Cloud structure ready
func init() {

	cache = goaws.GetCache()
	cache.Debugf(" cache -> %+v", cache)

	// First level goa sub commands
	GoaCmd.AddCommand(&RegionCmd)
	GoaCmd.AddCommand(&InstanceCmd)
	GoaCmd.AddCommand(&SnapshotCmd)
	GoaCmd.AddCommand(&VolumeCmd)

	// go instance terminate iid iid2 iid3 ...
	InstanceCmd.AddCommand(&TerminateInstancesCmd)

	// Second level volume commands
	VolumeCmd.AddCommand(&VolumeDeleteCmd)
}

// Execute the RootCommand
func Execute() {
	if err := GoaCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// GoaDo runs the root command
func cmdGoa(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to Goa! ")
	fmt.Println("  cache ", cache)
}

// list all regions
func cmdRegions(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("expected (regions) got ()")
	}
	fmt.Printf("Regions[%d]: \n", len(regions))
	fmt.Printf("%s", strings.Join(regions, "\n"))
	fmt.Printf("\n")
}

// List the instances
func cmdInstances(cmd *cobra.Command, args []string) {
	var regions []string
	if regions = goaws.Regions(); regions == nil {
		log.Fatal("  expected (regions) got ()")
	}

	for _, region := range regions {
		fmt.Printf(" getting region %s ... \n ", region)
		cl := goaws.GetCloud(region)
		cl.Instmap = goaws.GetInstances(region)
		for _, inst := range cl.Instmap {
			fmt.Printf("%s %s %s %s\n", inst.InstanceId, inst.VolumeId, inst.State.Name, inst.Region)
		}
	}
}

// List the instances
func cmdTerminateInstances(cmd *cobra.Command, args []string) {
	if err := goaws.TerminateInstances(args[0], args[1:]...); err != nil {
		log.Errorf("  problems with Terminate Instances %v", err)
		return
	}
}

// List volumes - TODO: check the regions argument
func cmdVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	var volumes goaws.Volmap
	for _, region := range regions {
		// See if the cache is working
		if volumes = goaws.GetVolumes(region); volumes == nil {
			log.Warning("  failed to get volumes from AWS ", region)
		}
		for _, vol := range volumes {
			fmt.Println(vol)
		}
	}
}

// Delete Volume specified by vol-xxxxxx arg[0]
func cmdDeleteVolume(cmd *cobra.Command, args []string) {
	var vol *goaws.Volume
	volid := args[0]

	if vol = goaws.GetVolume(volid); vol == nil {
		log.Errorf("  failed to find volume of %s", volid)
		return
	}
	if err := goaws.DeleteVolume(vol.VolumeId); err != nil {
		log.Errorf("  failed deleting volume %+v", err)
		return
	}
	fmt.Printf("  volume: %+v\n", vol)
}

// List Snapshots
func cmdSnapshots(cmd *cobra.Command, args []string) {
	panic("todo impletement snapshot list ")
}

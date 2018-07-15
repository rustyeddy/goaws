/*
Goa commands aws management utilities.
*/
package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	cache *goaws.Cache // store.Store

	// GoaCmd is the root command
	goaCmd = cobra.Command{
		Use:   "goa",
		Short: "Manage AWS Instances and Volumes",
		Run:   cmdGoa,
	}

	// RegionsCmd list regions
	regionCmd = cobra.Command{
		Use:   "region",
		Short: "List AWS Regions",
		Run:   cmdRegions,
	}

	// SnapCmd list snap shots
	snapshotCmd = cobra.Command{
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
	goaCmd.AddCommand(&regionCmd)
	goaCmd.AddCommand(&instanceCmd)
	goaCmd.AddCommand(&snapshotCmd)
	goaCmd.AddCommand(&volumeCmd)

	// Second level volume commands
	volumeCmd.AddCommand(&volumeDeleteCmd)
}

// Execute the RootCommand
func Execute() {
	if err := goaCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// GoaDo runs the root command
func cmdGoa(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to Goa! ")
	fmt.Println("        cache ", cache)
}

// List Snapshots
func cmdSnapshots(cmd *cobra.Command, args []string) {
	panic("todo impletement snapshot list ")
}

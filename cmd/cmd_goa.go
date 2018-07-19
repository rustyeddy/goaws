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
	// GoaCmd is the root command
	goaCmd = cobra.Command{
		Use:     "goa",
		Short:   "Manage AWS Instances and Volumes",
		Run:     cmdGoa,
		Version: "2018-07-12",
	}
)

// Get the AWS Cloud structure ready
func init() {
	goaws.SetRegion("us-west-2") // just in case

	// First level goa sub commands
	goaCmd.AddCommand(&regionCmd)
	goaCmd.AddCommand(&instanceCmd)
	goaCmd.AddCommand(&volumeCmd)
	goaCmd.AddCommand(&s3Cmd)
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
	fmt.Println("\tversion", cmd.Version)
}

/*
Goa commands aws management utilities.

- Goa

*/
package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
	"github.com/spf13/cobra"
)

var (
	S       *store.Store
	C       = goaws.C
	RootCmd = cobra.Command{
		Use:   "goa --help",
		Short: "Manage AWS Instances and Volumes",
		Long:  `Gather AWS Volume and Image information, do stuff with it`,
		Run:   GoaDo,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	pflags := RootCmd.PersistentFlags()
	pflags.StringVarP(&C.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&C.Region, "region", "r", "", "Select region defaults to all")

	// Log related flags
	pflags.StringVarP(&C.Loglevel, "level", "L", "debug", "Select level of logging")
	pflags.StringVarP(&C.Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&C.Logfile, "logfile", "l", "stdout", "Set the logging output")
}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {

	if S == nil {
		S = store.UseStore("/srv/store")
	}

	goaws.LogConfig(map[string]string{
		"level":  C.Loglevel,
		"format": C.Logformat,
		"log":    C.Logfile,
	})
}

// Execute from the RootCommand
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// RunRoot runs the root command
func GoaDo(cmd *cobra.Command, args []string) {
	fmt.Println("Welcome to Goa! ")
	fmt.Println("  basedir  ", C.Basedir)
	fmt.Printf("  store %s\n", S.String())
}

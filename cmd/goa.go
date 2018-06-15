/*
Goa commands aws management utilities.

- Goa

*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	C      = goaws.C
	GoaCmd = cobra.Command{
		Use:   "goa --help",
		Short: "Manage AWS Instances and Volumes",
		Long:  `Gather AWS Volume and Image information, do stuff with it`,
		Run:   GoaDo,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	pflags := GoaCmd.PersistentFlags()
	pflags.StringVarP(&C.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&C.Region, "region", "r", "", "Select region defaults to all")

	// Log related flags
	pflags.StringVarP(&C.Loglevel, "level", "l", "debug", "Select level of logging")
	pflags.StringVarP(&C.Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&C.Logfile, "logfile", "F", "stdout", "Set the logging output")
}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {
	goaws.InitConfig()
}

// Execute from the RootCommand
func Execute() {
	if err := GoaCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// RunRoot runs the root command
func GoaDo(cmd *cobra.Command, args []string) {
	fmt.Println(" Goa for a show: ")
	fmt.Printf("     goa %s\n", strings.Join(os.Args[1:], " "))

	log.Error("This is an error message")
	log.Warn("This is a warning")
	log.Info("THis is info")
	log.Debug("This is debugging at its best")
}

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
	GoaCmd cobra.Command
	c      *goaws.Configuration
)

func init() {
	c = goaws.C

	GoaCmd = cobra.Command{
		Use:   "goa --help",
		Short: "Manage AWS Instances and Volumes",
		Long:  `Gather AWS Volume and Image information, do stuff with it`,
		Run:   Goa,
	}
}

// Setup the flags
func initConfig() {
	pflags := GoaCmd.PersistentFlags()
	pflags.StringVarP(&c.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&c.Region, "region", "r", "", "Select region defaults to all")

	// Log related flags
	pflags.StringVarP(&c.Loglevel, "level", "l", "debug", "Select level of logging")
	pflags.StringVarP(&c.Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&c.Logfile, "logfile", "F", "stdout", "Set the logging output")
}

// Execute from the RootCommand
func Execute() {
	if err := GoaCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// RunRoot runs the root command
func Goa(cmd *cobra.Command, args []string) {
	fmt.Println(" Goa for a show")
	fmt.Printf("     args -- %s", strings.Join(os.Args[1:], ", "))
}

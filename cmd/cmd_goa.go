/*
Goa commands aws management utilities.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

// Configuration global to the goa command
type Configuration struct {
	Basedir string   // basedir for cache
	Region  string   // current region
	Regions []string // available regions
	Cfgfile string
	Verbose bool

	Logfile   string
	Logformat string
	Loglevel  string
}

var (
	// GoaCmd is the root command
	goaCmd = cobra.Command{
		Use:     "goa",
		Short:   "Manage AWS Instances and Volumes",
		Run:     cmdGoa,
		Version: "2018-07-12",
	}

	Config, FileConfig Configuration
)

// Get the AWS Cloud structure ready
func init() {
	goaws.SetRegion("us-west-2") // just in case
	cobra.OnInitialize(initConfig)

	pflags := goaCmd.PersistentFlags()
	pflags.StringVarP(&Config.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&Config.Region, "region", "r", "", "Select region defaults to all")
	pflags.StringVarP(&Config.Cfgfile, "cfgfile", "c", ".config/goa.json", "Specify the configuration file")
	pflags.BoolVarP(&Config.Verbose, "verbose", "v", false, "Get more info on output")

	// Log related flags
	pflags.StringVarP(&Config.Loglevel, "level", "L", "debug", "Select level of logging")
	pflags.StringVarP(&Config.Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&Config.Logfile, "logfile", "l", "stdout", "Set the logging output")

	// First level goa sub commands
	goaCmd.AddCommand(&regionCmd)
	goaCmd.AddCommand(&instanceCmd)
	goaCmd.AddCommand(&volumeCmd)
	goaCmd.AddCommand(&s3Cmd)

}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Can not find homedir")
	}

	// Determine the config file name
	cfgfile := home + "/.config/goa.json"

	// Read in the config file
	jbuf, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		log.Fatalf(cfgfile, " ", err)
	}

	if err = json.Unmarshal(jbuf, &FileConfig); err != nil {
		log.Fatal(cfgfile, " ", err)
	}

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

package cmd

import (
	"encoding/json"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/rustyeddy/logrus"

	"github.com/spf13/cobra"
)

// Configuration global to the goa command
type Configuration struct {
	Basedir string   // basedir for cache
	Region  string   // current region
	Regions []string // available regions
	Cfgfile string

	Logfile   string
	Logformat string
	Loglevel  string
}

var (
	Config, FileConfig, EnvConfig Configuration
)

func init() {
	cobra.OnInitialize(initConfig)

	pflags := goaCmd.PersistentFlags()
	pflags.StringVarP(&Config.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&Config.Region, "region", "r", "", "Select region defaults to all")
	pflags.StringVarP(&Config.Cfgfile, "cfgfile", "c", ".config/goa.json", "Specify the configuration file")

	// Log related flags
	pflags.StringVarP(&Config.Loglevel, "level", "L", "debug", "Select level of logging")
	pflags.StringVarP(&Config.Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&Config.Logfile, "logfile", "l", "stdout", "Set the logging output")
}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {
	log.Debugln("~~> initConfig entered")
	defer log.Debugln("<~~ initConfig existing ...")

	// TODO get the config file in there ...
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

	// TODO Merge FileConfig into regular config and read ENV
}

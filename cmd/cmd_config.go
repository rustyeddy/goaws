package cmd

import (
	"encoding/json"
	"io/ioutil"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/rustyeddy/logrus"

	"github.com/spf13/cobra"
)

var (
	Basedir string   // basedir for cache
	Region  string   // current region
	Regions []string // available regions

	Logfile   string
	Logformat string
	Loglevel  string
)

func init() {
	cobra.OnInitialize(initConfig)
	pflags := RootCmd.PersistentFlags()
	pflags.StringVarP(&Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&Region, "region", "r", "", "Select region defaults to all")

	// Log related flags
	pflags.StringVarP(&Loglevel, "level", "L", "debug", "Select level of logging")
	pflags.StringVarP(&Logformat, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&Logfile, "logfile", "l", "stdout", "Set the logging output")
}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {
	log.Debug("~~> initConfig entered")
	defer log.Debug("<~~ initConfig existing ...")

	// TODO get the config file in there ...
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Can not find homedir")
	}

	cfgfile := home + "/.config/goa.json"
	jbuf, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		log.Fatalf(cfgfile, " ", err)
	}
	log.Fatalf("%v", string(jbuf[0:90]))

	var config map[string]string
	if err = json.Unmarshal(jbuf, &config); err != nil {
		log.Fatal(cfgfile, " ", err)
	}
}

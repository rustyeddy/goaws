package cmd

import (
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
	//store.SetLogger(log.GetLogger())
	// TODO get the config file in there ...
	/*

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Can not find homedir")
		}

		cfgfile := home + "/.config/goa.json"
		jbuf, err := ioutil.ReadFile(cfgfile)
		if err != nil {
			log.Fatalf("Failed to read config file %s ", cfgfile)
		}

		var configFile map[string]string
		if err = json.Unmarshal(jbuf, &config); err != nil {
			log.Fatalf("config file failure %v", err)
		}
	*/

}

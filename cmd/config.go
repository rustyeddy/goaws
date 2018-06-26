package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/rustyeddy/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigLogger defines the behavior of the logger
type ConfigLogger struct {
	Level   string
	Logfile string
	Format  string
}

// Config is other config stuff
type Config struct {
	Basedir string
	Region  string
	Regions []string

	ConfigFile string

	ConfigLogger
}

var (
	config Config
)

func init() {
	cobra.OnInitialize(initConfig)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")

	viper.AddConfigPath("$HOME/.goa")         // call multiple times to add many search paths
	viper.AddConfigPath("$HOME/.config/goa/") // path to look for the config file in
	viper.AddConfigPath("/etc/goa/")          // path to look for the config file in
	viper.AddConfigPath(".")                  // optionally look for config in the working directory

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})

	pflags := RootCmd.PersistentFlags()
	pflags.StringVarP(&config.Basedir, "dir", "d", "/srv/goaws/", "base project directory")
	pflags.StringVarP(&config.Region, "region", "r", "", "Select region defaults to all")
	viper.BindPFlag("basedir", RootCmd.PersistentFlags().Lookup("basedir"))
	viper.BindPFlag("region", RootCmd.PersistentFlags().Lookup("region"))
	viper.SetDefault("basedir", "/srv/goaws/store")

	// Log related flags
	pflags.StringVarP(&config.Level, "level", "L", "debug", "Select level of logging")
	pflags.StringVarP(&config.Format, "format", "f", "json", "Select level of logging")
	pflags.StringVarP(&config.Logfile, "logfile", "l", "stdout", "Set the logging output")
	viper.BindPFlag("level", RootCmd.PersistentFlags().Lookup("level"))
	viper.BindPFlag("format", RootCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("out", RootCmd.PersistentFlags().Lookup("out"))
	viper.SetDefault("level", "warn")
	viper.SetDefault("format", "json")
	viper.SetDefault("out", "out")

	// Read the configuration
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// initConfig run everytime the subcmcd Execute() is run
func initConfig() {
	log.Debug("~~> initConfig entered")
	defer log.Debug("<~~ initConfig existing ...")

	// Don't forget to read config either from cfgFile or from home directory!
	if config.ConfigFile != "" {
		viper.SetConfigFile(config.ConfigFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Can not find homedir")
		}

		// Search config in home directory
		viper.AddConfigPath(home + "/.config/goa")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

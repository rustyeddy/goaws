package goaws

import log "github.com/rustyeddy/logrus"

type Configuration struct {
	Basedir string   // basedir specifically for store
	Regions []string // regions we care about
	Region  string   // region currently being worked on

	Loglevel  string // string representation of
	Logformat string // json or txt
	Logfile   string // file to log to

	Cache bool // should we cache what we get?
}

var (
	DefaultConfig Configuration
)

func init() {
	DefaultConfig = Configuration{
		Basedir: C.Basedir, // were we live on the filesystem
		Regions: nil,       // regions we care about
		Region:  "",        // Current region

		Loglevel:  "debug",     // debug, info, warn, error, fatal, panic
		Logformat: "text",      // json, text
		Logfile:   "goaws.log", // hopefully set in aws
		Cache:     true,        // caching is on by default
	}
	log.Debugln("leaving config init")
}

// SaveConfig will save the running configuration to the files specified
// if no file is provided the configuraion file will be written to the
// default location.
func (cfg *Configuration) SaveConfig() error {
	_, err := S.StoreObject("config", C)
	return err
}

// ReadConfig will restore the config back into ourselves.
func (cfg *Configuration) ReadConfig() error {
	err := S.FetchObject("config", &C)
	return err
}

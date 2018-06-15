package goaws

type Configuration struct {
	Basedir   string   // basedir specifically for store
	Regions   []string // regions we care about
	Region    string   // region currently being worked on
	Loglevel  string   // string representation of
	Logformat string   // json or txt
	Logfile   string   // file to log to
}

var (
	DefaultConfig Configuration
)

func init() {
	DefaultConfig = Configuration{
		Basedir:   C.Basedir, // were we live on the filesystem
		Regions:   nil,       // regions we care about
		Region:    "",        // Current region
		Loglevel:  "debug",   // debug, info, warn, error, fatal, panic
		Logformat: "json",    // json, text
	}
}

func InitConfig() {
	// Probably need to read some command line arguments to
	// see what needs to be done
}

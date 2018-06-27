package goaws

// LogConfig holds configuration for Logging
type LogConfig struct {
	Level   string
	Logfile string
	Format  string
}

// Config holds variables that may be configured
type Config struct {
	Name       string
	Basedir    string
	Region     string
	Configfile string
	LogConfig
}

var (
	configDebug Config // Possible configs for appropriate env
	configTest  Config // Test local storage
	configProd  Config // Production enviroments
	config      Config // Current / Running config
)

func init() {
	configDebug = Config{
		Name:      "debug",
		Basedir:   "/srv/goaws/cache/",
		Region:    "us-west-2",
		LogConfig: LogConfig{Level: "debug", Logfile: "stdout", Format: "text"},
	}
	configTest = Config{
		Name:      "test",
		Basedir:   "test/",
		Region:    "us-west-2",
		LogConfig: LogConfig{Level: "info", Logfile: "goa-test", Format: "json"},
	}
	configProd = Config{
		Name:      "test",
		Basedir:   "test/",
		Region:    "us-west-2",
		LogConfig: LogConfig{Level: "warn", Logfile: "goa-prod", Format: "json"},
	}
	config = configDebug
}

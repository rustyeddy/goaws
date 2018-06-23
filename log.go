package goaws

import (
	"os"

	log "github.com/rustyeddy/logrus"
)

func init() {
	// Set some resonable defaults
	log.SetFormatter(&log.TextFormatter{}) // JSON formatted output
	log.SetLevel(log.DebugLevel)           // High default log level
	log.SetOutput(os.Stdout)               // print to stdout by default
}

// LogConfig will accept a map of key value strings that can
// optionally set logging level, format and output.
func LogConfig(cfg map[string]string) {
	for n, v := range cfg {
		switch n {
		case "level":
			setLevelString(v)
		case "format":
			setFormatString(v)
		case "log":
			setLogFilename(v)
		default:
			log.Warn("LogConfig bad cmd ", n)
		}
	}
}

func setLevelString(lstr string) {
	var lvl map[string]log.Level
	lvl = make(map[string]log.Level)

	lvl["debug"] = log.DebugLevel
	lvl["info"] = log.InfoLevel
	lvl["warn"] = log.WarnLevel
	lvl["error"] = log.ErrorLevel
	lvl["fatal"] = log.FatalLevel
	lvl["panic"] = log.PanicLevel

	if l, e := lvl[lstr]; e {
		log.SetLevel(l)
	}
}

func setFormatString(fstr string) {
	switch fstr {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.Warning("unknown format string: ", fstr)
	}
}

func setLogFilename(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Errorf("setLogFname: %s -> %v", fname, err)
	} else {
		log.SetOutput(f)
	}
}

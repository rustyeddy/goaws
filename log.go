package goaws

import (
	"os"

	log "github.com/rustyeddy/logrus"
)

func init() {
	// Set some resonable defaults
	log.SetFormatter(&log.TextFormatter{}) // JSON formatted output
	log.SetLevel(log.WarnLevel)            // High default log level
	log.SetOutput(os.Stdout)               // print to stdout by default
}

// SetLogConfig will accept a map of key value strings that can
// optionally set logging level, format and output.
func SetLogConfig(cfg map[string]string) {
	for n, v := range cfg {
		switch n {
		case "level":
			setLevelString(v)
		case "format":
			setFormatString(v)
		case "log":
			setLogFilename(v)
		default:
			log.Warn("LogConfig unknown command ", n)
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
		log.Warning("Unknown format string: ", fstr)
	}
}

// logfile name or the string "stdout" for os.Stdout
func setLogFilename(fname string) {
	f := os.Stdout
	switch fname {
	case "stdout":
		f = os.Stdout
	case "stderr":
		f = os.Stderr
	default:
		var err error
		if f, err = os.Open(fname); err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"file":  fname,
			}).Fatalf("failed to set log file, will not contine")
		} else {
			log.SetOutput(f)
		}
	}
}

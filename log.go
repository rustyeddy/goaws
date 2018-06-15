package goaws

import (
	"os"

	log "github.com/rustyeddy/logrus"
)

func init() {

	// Set the logging defaults
	log.SetFormatter(&log.JSONFormatter{}) // JSON formatted output
	log.SetLevel(log.DebugLevel)           // High default log level
	fname := "goaws.log"
	f, err := os.Open(fname)
	if err != nil {
		log.SetOutput(f)
	}
}

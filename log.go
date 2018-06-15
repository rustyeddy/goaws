package aws

import (
	"os"

	log "github.com/rustyeddy/logrus"
)

// Setup the logger to log JSON to goaws.log at LevelDebug
func initLog() {

	// Setup the logger (logrus). Set JSON as the output format
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	fname := "goaws.log"
	f, err := os.Open(fname)
	if err != nil {
		log.SetOutput(f)
	}
}

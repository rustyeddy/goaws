/*
This AWS utily scans and indexes all regions for Instances and
volumes.  The instances and volumes can be managed from these
indexes including deleting them.
*/
package main

import (
	"flag"

	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
)

// ======================================================================

var (
	fetch    = flag.Bool("fetch", false, "fetch from AWS, default false, read cached files")
	loglevel = flag.String("loglevel", "", "Errors and above always logged debug and info")
	region   = flag.String("region", "", "pick the region you want to run this command on")
	pattern  = flag.String("pattern", "/srv/goaws/*/*", "Glob pattern to match files")

	deleteIds []string // String of vol-* or i-* ids to delete

	// Store for
	s *store.Store
)

func init() {
	logInit()
	s = store.New("goaws")
}

func main() {
	flag.Parse()

	// Change logging levels?
	switch *loglevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	default:
		log.Error("unknown loglevel ", *loglevel)
	}

	// Get inventory from region if it has been provided
	var inv *Inventory
	if *region != "" {
		inv = GetInventory(*region)
	}

	// Now figure out what we need to do and do it.
	switch {
	case inv == nil && *fetch:
		FetchInventories()
	case inv == nil && *fetch == false:
		ReadInventories()
	case inv != nil && *fetch:
		inv.FetchInventory()
	case inv != nil && *fetch == false:
		inv.ReadFiles(FindFiles("test/*-" + inv.Name + ".json"))
	default:
		log.Error("we hit a default bummer")
	}
}

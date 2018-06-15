/*
This AWS utily scans and indexes all regions for Instances and
volumes.  The instances and volumes can be managed from these
indexes including deleting them.
*/
package main

import (
	"flag"
	"os"

	"github.com/rustyeddy/store"
)

var (
	fetch    = flag.Bool("fetch", false, "fetch from AWS, default false, read cached files")
	loglevel = flag.String("loglevel", "", "Errors and above always logged debug and info")
	region   = flag.String("region", "", "pick the region you want to run this command on")
	pattern  = flag.String("pattern", "/srv/goaws/*/*", "Glob pattern to match files")

	// Store for
	s *store.Store
)

func main() {
	flag.Parse()
	RootCmd.Run(os.Args())
}

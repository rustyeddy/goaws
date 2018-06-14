/*
This AWS utily scans and indexes all regions for Instances and
volumes.  The instances and volumes can be managed from these
indexes including deleting them.
*/
package main

import (
	"flag"
)

// ======================================================================

var (
	fetch   bool
	verbose bool

	delete []string // String of vol-* or i-* ids to delete
)

func init() {
	flag.BoolVar(&fetch, "fetch", false, "fetch from AWS, default false, read cached files")
	flag.BoolVar(&verbose, "verbose", false, "tobble verbosity")
}

func main() {
	flag.Parse()
	if fetch {
		FetchInventories()
	}
}

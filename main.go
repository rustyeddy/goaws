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
	fetch        bool
	verbose      bool

	delete    []string // String of vol-* or i-* ids to delete
)

func init() {
	flag.BoolVar(&fetch, "fetch", false, "fetch from AWS, default false, read cached files")
	flag.BoolVar(&verbose, "verbose", false, "tobble verbosity")
	flag.BoolVar(&printRegions, "list-regions", false, "print the regions")

	flag.
}

type output struct {
	buf string
	fmt string
	err error
}

func main() {
	flag.Parse()



	out := ""
	if printRegions {
		out = PrintRegions()
	} else if fetch {
		out = FetchInventories()
	} else {
		out = ReadFiles(flag.Args())
	}

}

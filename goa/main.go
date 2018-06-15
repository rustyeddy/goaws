/*
This AWS utily scans and indexes all regions for Instances and
volumes.  The instances and volumes can be managed from these
indexes including deleting them.
*/
package main

import (
	"github.com/rustyeddy/goaws/cmd"
)

func main() {
	cmd.Execute()
}

package cmd

import (
	"fmt"
	"strings"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

// list all regions
func cmdRegions(cmd *cobra.Command, args []string) {

	var regions []string
	var err error
	if regions, err = goaws.Regions(); err != nil {
		log.Fatal("expected (regions) got (%v)", err)
	}
	fmt.Printf("Regions[%d]: \n", len(regions))
	fmt.Printf("\n%s", strings.Join(regions, "\n"))
	fmt.Printf("\n")
}

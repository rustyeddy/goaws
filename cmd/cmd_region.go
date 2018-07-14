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
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("expected (regions) got ()")
	}
	fmt.Printf("Regions[%d]: \n", len(regions))
	fmt.Printf("%s", strings.Join(regions, "\n"))
	fmt.Printf("\n")
}

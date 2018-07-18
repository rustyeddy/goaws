package cmd

import (
	"fmt"

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
	icount := 0
	for _, region := range regions {
		fmt.Printf("  %s ", region)
		if insts := goaws.Instances(region); insts != nil {
			icount = len(insts)
		} else {
			icount = 0
		}
		fmt.Printf("  %s %d instances \n", region, icount)
	}
	fmt.Printf("done.")
	fmt.Printf("\n")
}

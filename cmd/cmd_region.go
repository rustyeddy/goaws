package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	// RegionsCmd list regions
	regionCmd = cobra.Command{
		Use:   "region",
		Short: "List AWS Regions",
		Run:   cmdRegions,
	}
)

// list all regions
func cmdRegions(cmd *cobra.Command, args []string) {

	var regions []string
	var err error
	if regions = goaws.Regions(); regions == nil {
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

// GetRegions will return a list of regions the corresponding
// command will be applied to.  The command will determine how the
// default region is selected, in general it will be the default
// region as defined in the aws configuration.
func GetRegions(cmd *cobra.Command, args []string) (regions []string) {

	if Config.Region != "" {
		regions = []string{Config.Region}
	} else {
		if regions = goaws.Regions(); regions == nil {
			log.Fatal("  expected (regions) got (nil)")
		}
	}
	return regions
}

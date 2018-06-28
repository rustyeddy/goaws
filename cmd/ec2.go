package cmd

import (
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	regions []string
	ec2Cmd  = cobra.Command{
		Use:   "instances",
		Short: "list EC2 instances",
		Run:   DoEC2,
	}
)

func init() {
	RootCmd.AddCommand(&ec2Cmd)
}

func DoEC2(cmd *cobra.Command, args []string) {

	// Walk the regions getting inventories
	if regions = goaws.Regions(); regions == nil {
		log.Warn("Failed to grab a list of regions... ")
		return
	}
	var inv *goaws.Inventory
	for _, n := range regions {
		if inv = goaws.GetInventory(n); inv == nil {
			log.Errorf("failed to get inventory (DoEC2) %s ", n)
		}
		log.Printf("  got an inventory %+v with no data ", inv)
	}
}

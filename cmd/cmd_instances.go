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

// DoEC2 executes the EC2
func DoEC2(cmd *cobra.Command, args []string) {

	regions := goaws.Regions()
	if regions == nil {
		log.Fatalf("  expected list of regions got ()")
	}

	for _, r := range regions {
		instances := goaws.FetchInstances(r)
		if instances == nil {
			log.Errorf("  no instances for region %s", r)
			continue
		}
		// Print some instances information

	}
}

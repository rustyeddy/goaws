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
	inv := goaws.FetchInventories()
	if inv == nil {
		log.Fatalf("  inventories %+v ", inv)
	}
	log.Printf(" inventories %+v ", inv)
}

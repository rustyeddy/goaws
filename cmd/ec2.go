package cmd

import (
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	ec2Cmd = cobra.Command{
		Use:   "instances",
		Short: "list EC2 instances",
		Run:   DoEC2,
	}
)

func init() {
	RootCmd.AddCommand(&ec2Cmd)
}

func DoEC2(cmd *cobra.Command, args []string) {
	log.Printf("DoEc2")
}

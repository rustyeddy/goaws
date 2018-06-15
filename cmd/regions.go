package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	regionCmd = cobra.Command{
		Use:   "region cmd ...",
		Short: "manage regions with this wonderful cli",
		Long:  "AWS divides most things into regions, manage them",
		Run:   regionDo,
	}

	regionListCmd = cobra.Command{
		Use:   "ls ...",
		Short: "list regions",
		Long:  "list available regions",
		Run:   regionListDo,
	}
)

func init() {
	GoaCmd.AddCommand(&regionCmd)
	regionCmd.AddCommand(&regionListCmd)
}

func regionDo(cmd *cobra.Command, args []string) {
	log.Debug("Fetching regions")

	regs := goaws.Regions()
	if regs != nil {
		log.Error("failed to get regions")
		return
	}
	fmt.Println(regs)
}

func regionListDo(cmd *cobra.Command, args []string) {
	log.Debug("Region list")
	regions := goaws.Regions()
	if regions == nil {
		log.Error("failed get AWS regions")
	}
}

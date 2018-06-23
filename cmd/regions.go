package cmd

import (
	"fmt"
	"strings"

	"github.com/rustyeddy/goaws"
	"github.com/spf13/cobra"

	log "github.com/rustyeddy/logrus"
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
	RootCmd.AddCommand(&regionCmd)
	regionCmd.AddCommand(&regionListCmd)
}

func regionDo(cmd *cobra.Command, args []string) {
	log.Debug("Fetching regions")
	regs := goaws.Regions()
	if regs != nil {
		log.Error("  failed to get regions")
		return
	}
	fmt.Println(regs)
}

func regionListDo(cmd *cobra.Command, args []string) {

	var regions []string
	if regions = goaws.Regions(); regions == nil {
		// I got nothing to say, no regions have been found
	}
	log.Println(strings.Join(regions, "\n"))
}

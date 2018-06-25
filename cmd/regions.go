package cmd

import (
	"strings"

	"github.com/rustyeddy/goaws"
	"github.com/spf13/cobra"

	log "github.com/rustyeddy/logrus"
)

var (
	regions   []string
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

	log.Debugln("region do ..")
	log.Debugln("  --> check cache for region list ")
	defer log.Debugln("Regions[%d]: %+v", len(regions), regions)

	regions := goaws.Regions()
	if regions == nil {
		log.Error("Unable to file any regions, dieing ")
	}
}

func regionListDo(cmd *cobra.Command, args []string) {
	var regions []string
	log.Debugln("  region LIST Do ..")
	if regions = goaws.Regions(); regions == nil {
		// I got nothing to say, no regions have been found
		log.Println("  store empty of objects ")
	}
	log.Println(strings.Join(regions, "\n"))
}

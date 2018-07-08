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
		Use:   "regions",
		Short: "manage regions with this wonderful cli",
		Long:  "AWS divides most things into regions, manage them",
		Run:   regionDo,
	}
)

func init() {
	RootCmd.AddCommand(&regionCmd)
}

func regionDo(cmd *cobra.Command, args []string) {

	fmt.Println("list all regions ..")
	regions := goaws.Regions()
	if regions == nil {
		log.Error("  # Unable to file any regions, dieing ")
	}

	fmt.Printf("Regions[%d]: \n", len(regions))
	fmt.Printf("%s", strings.Join(regions, "\n"))
	fmt.Printf("\n")
}

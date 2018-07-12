package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdVolumes = cobra.Command{
		Use:   "volumes",
		Short: "list volumes",
		Run:   doVolumes,
	}
	cmdVol = cobra.Command{
		Use:   "vol",
		Short: "Get information or content of a volume",
		Run:   doVol,
	}
)

func init() {
	RootCmd.AddCommand(&cmdVolumes)
	RootCmd.AddCommand(&cmdVol)
}

func doVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	var volumes goaws.Volmap
	for _, region := range regions {
		// See if the cache is working
		fmt.Println("doVolumes is calling get volumes ")
		if volumes = goaws.GetVolumes(region); volumes == nil {
			log.Fatal("  failed to get volumes from AWS")
		}
		for _, vol := range volumes {
			fmt.Println(vol)
		}
	}
}

func doVol(cmd *cobra.Command, args []string) {
	vid := args[0]
	vol := goaws.GetVolume(vid)
	if vol == nil {
		log.Errorf("  failed to find volume of %s", vid)
		return
	}
	fmt.Printf("  volume: %+v\n", vol)
}

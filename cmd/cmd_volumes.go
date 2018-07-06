package cmd

import (
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	volumeCommand = cobra.Command{
		Use:   "volumes",
		Short: "list volumes",
		Run:   doVolumes,
	}
)

func init() {
	RootCmd.AddCommand(&volumeCommand)
}

func doVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	var volumes []goaws.VDisk
	for _, region := range regions {
		// See if the cache is working
		if volumes = goaws.GetVolumes(region); volumes == nil {
			log.Fatal("  failed to get volumes from AWS")
		}
		log.Fatalf("  volumes %+v \n\n")
	}
}

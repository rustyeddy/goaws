package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

// List volumes - TODO: check the regions argument
func cmdVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	var volumes goaws.Volmap
	for _, region := range regions {
		// See if the cache is working
		if volumes = goaws.GetVolumes(region); volumes == nil {
			log.Warning("  failed to get volumes from AWS ", region)
		}
		for _, vol := range volumes {
			fmt.Println(vol)
		}
	}
}

// Delete Volume specified by vol-xxxxxx arg[0]
func cmdDeleteVolume(cmd *cobra.Command, args []string) {
	var vol *goaws.Volume
	volid := args[0]

	if vol = goaws.GetVolume(volid); vol == nil {
		log.Errorf("  failed to find volume of %s", volid)
		return
	}
	if err := goaws.DeleteVolume(vol.VolumeId); err != nil {
		log.Errorf("  failed deleting volume %+v", err)
		return
	}
	fmt.Printf("  volume: %+v\n", vol)
}

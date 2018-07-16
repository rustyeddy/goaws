package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	// VolCmd list Volumes from all regions (or the -region flag)
	volumeCmd = cobra.Command{
		Use:     "volume",
		Short:   "Manage AWS Volumes",
		Aliases: []string{"vol", "volumes"},
		Run:     cmdVolumes,
	}

	volumeDeleteCmd = cobra.Command{
		Use:     "delete",
		Aliases: []string{"del", "rm"},
		Short:   "delete the volume by vol-Id",
		Run:     cmdDeleteVolume,
	}
)

func init() {
	volumeCmd.AddCommand(&volumeDeleteCmd)
}

// List Volumes - TODO: check the regions argument
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
	var (
		vol   *goaws.Volume
		cl    *goaws.AWSCloud
		volid string = args[0]
	)
	cl = goaws.GetCloud(Config.Region)
	if err := cl.DeleteVolume(volid); err != nil {
		log.Errorf("  failed deleting volume %+v", err)
		return
	}
	fmt.Printf("  volume: %+v\n", vol)
}

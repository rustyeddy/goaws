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
	var regions []string
	if len(args) > 0 {
		regions = append(regions, args[0:]...)
	} else {
		regions = goaws.Regions()
	}

	for _, region := range regions {
		fmt.Printf("Volumes for region %s ", region)
		volumes := goaws.Volumes(region)
		if volumes == nil {
			fmt.Println("[0]")
			continue
		}
		fmt.Println(len(volumes))
		for vid, vol := range volumes {
			fmt.Printf("  %s %+v\n", vid, vol.State)
		}
	}
	fmt.Println("done ... ")
}

// Delete Volume specified by vol-xxxxxx arg[0]
func cmdDeleteVolume(cmd *cobra.Command, args []string) {
	region := args[0]
	volid := args[1]
	if err := goaws.DeleteVolume(region, volid); err != nil {
		log.Errorf("  failed deleting volume %+v", err)
		return
	}
	fmt.Printf("  volume: %s\n", volid)
}

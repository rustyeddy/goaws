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
	output := ""
	volumes := goaws.Volumes(Config.Region)
	for vid, _ := range volumes {
		output += vid + "\n"
	}
	fmt.Println(output)
}

// Delete Volume specified by vol-xxxxxx arg[0]
func cmdDeleteVolume(cmd *cobra.Command, args []string) {
	volid := args[0]
	if err := goaws.DeleteVolume(Config.Region, volid); err != nil {
		log.Errorf("  failed deleting volume %+v", err)
		return
	}
	fmt.Printf("  volume: %s\n", volid)
}

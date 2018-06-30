package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	volmap        map[string]*ec2.CreateVolumeOutput
	volumeCommand = cobra.Command{
		Use:   "volumes",
		Short: "list volumes",
		Run:   doVolumes,
	}
)

func init() {
	RootCmd.AddCommand(&volumeCommand)
	volmap = make(map[string]*ec2.CreateVolumeOutput)
}

func doVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	cache := goaws.Cache()
	for _, region := range regions {
		volmap := goaws.FetchVolumes(region)
		if volmap == nil {
			log.Errorf("  no volumes for region %s ", region)
			continue
		}
		log.Fatalf(" volmap %+v ", volmap)
		log.Debugln("  save the result to volumes map for ", region)
	}

	if volmap != nil && len(volmap) > 0 {
		cache = goaws.Cache()
		if cache == nil {
			log.Error("  failed to get cache for goaws ")
			return
		}
		obj, err := cache.StoreObject("volume-map", &volmap)
		if err != nil {
			log.Errorf("  failed to store in cache %v ", err)
			return
		}
		log.Debugf("  store volume %s", obj.Path)
	}

	for n, vol := range volmap {
		fmt.Printf(" volumes %s -> %v", n, vol)
	}

}

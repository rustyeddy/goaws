package cmd

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	volumes       map[string]*ec2.CreateVolumeOutput
	volumeCommand = cobra.Command{
		Use:   "volumes",
		Short: "list volumes",
		Run:   doVolumes,
	}
)

func init() {
	RootCmd.AddCommand(&volumeCommand)
	volumes = make(map[string]*ec2.CreateVolumeOutput)
}

func doVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}

	for _, region := range regions {
		results, volmap := goaws.FetchVolumes(region)
		if results == nil {
			log.Errorf("  no volumes for region %s ", region)
			continue
		}
		if volmap == nil {
			log.Errorf("  no volume map for region %s ", region)
			continue
		}
		volmap[region] = results
		log.Debugln("  save the result to volumes map for ", region)

	}
	if volmap != nil && len(volmap) > 0 {
		cache := goaws.Cache()
		if cache == nil {
			log.Error("  failed to get cache for goaws ")
			return
		}
		obj, err := cache.StoreObject("volume-map", &volmap)
		if err != nil {
			log.Errorf("  failed to store in cache %v ", err)
			return
		}
	}

	if (volumes != nil) && (len(volumes) > 0) {
		fname := egion + "-volumes"
		obj, err := cache.StoreObject(fname, volumes)
		if err != nil {
			log.Errorf("  failed store a cache version of %s volumes %v", fname, err)
			return
		}
		log.Debugf("  stashed object at path %s ", obj.Path)
	}
	fmt.Printf(" volumes %s ", strings.Join(volumes, "\n\t"))
}

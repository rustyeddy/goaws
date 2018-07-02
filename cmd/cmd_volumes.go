package cmd

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

type Disk struct {
	VolumeId   string
	InstanceId string
	CreateTime string
	AZ         string
	Size       int
}

var (
	volmap        map[string]*ec2.DescribeVolumesOutput
	volumeCommand = cobra.Command{
		Use:   "volumes",
		Short: "list volumes",
		Run:   doVolumes,
	}
)

func init() {
	RootCmd.AddCommand(&volumeCommand)
	volmap = make(map[string]*ec2.DescribeVolumesOutput)
}

func doVolumes(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  failed to get the regions, can't continue ")
	}
	for _, region := range regions {
		volmap := goaws.FetchVolumes(region)
		if volmap == nil {
			log.Errorf("  no volumes for region %s ", region)
			continue
		}
		log.Debugln("  save the result to volumes map for ", region)

		//log.Fatalf(" %+v", volmap)

		//DisksFromVolume(volmap)
	}
}

// Return a map of disks
func DisksFromVolume(vout map[string]*ec2.CreateVolumeOutput) {

	/*
		mdisk := make(map[string]*Disk, 100)
		log.Fatalf(" %T ", vout)

			for region, volumes := range vout {
				for _, vol := range volumes {
					log.Fatalf(" vol: %T ", vol)
					mdisk[vol.VolumeId] = vol
				}
			}
	*/
	//return mdisk
}

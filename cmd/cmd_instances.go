package cmd

import (
	"fmt"

	"github.com/rustyeddy/goaws"
	log "github.com/rustyeddy/logrus"
	"github.com/spf13/cobra"
)

var (
	regions     []string
	cmdInstance = cobra.Command{
		Use:   "instances",
		Short: "list instances",
		Run:   doInstances,
	}

	cmdDeleteInstance = cobra.Command{
		Use:   "rminst",
		Short: "Remove an instance and associated volume(s) if any",
		Run:   doDeleteInstance,
	}
)

func init() {
	RootCmd.AddCommand(&cmdInstance)
	RootCmd.AddCommand(&cmdDeleteInstance)
}

// DoEC2 executes the EC2
func doInstances(cmd *cobra.Command, args []string) {
	regions := goaws.Regions()
	if regions == nil {
		log.Fatal("  expected list of regions got ()")
	}

	for _, region := range regions {
		cl := goaws.GetCloud(region)
		cl.Instmap = goaws.GetInstances(region)
		for iid, _ := range cl.Instmap {
			fmt.Println(iid)
		}
	}
}

// doDeleteInstance
func doDeleteInstance(cmd *cobra.Command, args []string) {
	// Find an instance and try to delete it..
	regions := goaws.Regions()
	for _, region := range regions {

		fmt.Printf("Try to delete something from %s\n ", region)
		cl := goaws.GetCloud(region)
		imap := goaws.GetInstances(region)
		for k, v := range imap {
			cl.Instmap[k] = v
		}

		for _, inst := range cl.Instmap {
			volid := inst.VolumeId
			if err := goaws.DeleteVolume(volid); err != nil {
				log.Fatalf("  failed to send DeleteVolume request %v", err)
			}
			//log.Fatalf("VOL:  %+v ", vol)
		}
	}
}

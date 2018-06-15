package aws

import (
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// FetchVolumes will retrieve instances from AWS, convert them to
// Go structures we can use, it also "caches" a version to the filesystem
func (inv *Inventory) FetchVolumes() {

	log.Println("GetVolumes from ", inv.Region)
	defer log.Println("  return GetVolumes ", inv.Region)

	if e := inv.GetEC2(); e != nil {
		req := e.DescribeVolumesRequest(&ec2.DescribeVolumesInput{})
		if result, err := req.Send(); err == nil {
			inv.indexVolumes(result.Volumes)
			inv.saveVolumes(result.Volumes)
		}
	}
}

func (inv *Inventory) saveVolumes(res []ec2.CreateVolumeOutput) {
	// Send the request to aws and wait for results
	if jbytes, err := json.Marshal(res); err == nil {
		fname := "etc/vol-" + inv.Region + ".json"
		if err = ioutil.WriteFile(fname, jbytes, 0644); err == nil {
			log.Debug("cached volumes in ", fname)
			return // our work is done
		}
	}
}

// DeleteVolume sends a request to delete the specified volume
func (inv *Inventory) DeleteVolume(volid string) {
	panic("TODO DELETE VOLUME")
}

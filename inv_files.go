package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

/*
  files.go: handles reading and writting AWS Specific JSON formatted
  files.  In otherwords we will read JSON encode files in an AWS
  API format.
*/

// FindFiles will gather up files, determine the region they are from
// Then process and index the files.
func FindFiles(pattern string) []string {
	paths, err := filepath.Glob(pattern)
	if err != nil {
		log.Error("failed to glob ", pattern, err)
		return nil
	}
	return paths
}

// ReadFiles will recieve the specified files into the inventory
func (inv *Inventory) ReadFiles(paths []string) {
	for _, p := range paths {
		inv.ReadFile(p)
	}
}

// ReadFile will injest the given file into the inventory
func (inv *Inventory) ReadFile(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("failed reading ", path)
	}
	id := path[0:7]
	switch id {
	case "etc/ins":
		inv.unmarshalInstances(buf)
	case "etc/vol":
		inv.unmarshalVolumes(buf)
	default:
		log.Fatal("expected (etc/vol) or (etc/inst) got path ", path)
	}
}

// Support functions for converting JSON -> Go
// ======================================================================

// unmarshalInstances from a byte slice (AWS) will also index the items
// read from file
func (inv *Inventory) unmarshalInstances(buf []byte) {
	log.Debugln("unmarshalling Volume buffer")
	var rlist []ec2.RunInstancesOutput
	err := json.Unmarshal(buf, &rlist)
	if err != nil {
		log.Error("instances ", err)
		return
	}

	// index these instances
	for _, rl := range rlist {
		for _, inst := range rl.Instances {
			inv.Instances[*inst.ImageId] = HostFromInstance(&inst)
		}
	}
}

// unmarshalVolumes from a buffer of bytes
func (inv *Inventory) unmarshalVolumes(buf []byte) []ec2.CreateVolumeOutput {
	log.Debugln("unmarshalling Volume buffer")
	var vols []ec2.CreateVolumeOutput
	err := json.Unmarshal(buf, &vols)
	if err != nil {
		log.Error("failed to unmarshal buffer ", string(buf[0:40]))
		return nil
	}

	// index the volumes we have read
	for _, vol := range vols {
		inv.Volumes[*vol.VolumeId] = DiskFromVolume(&vol)
	}
	return vols
}

package goaws

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// ReadInventories reads inventory files from files.  The region
// will be extracted by the filename.
func ReadInventories() {
	// First we need to locate the files we are going to read
	pattern := "cache/*/*.json"
	if currentRegion != "" {
		pattern = "cache/" + currentRegion + ".json"
	}

	paths := FindFiles(pattern)
	if paths == nil {
		log.Error("failed to read inventory files")
	}

	// Now we need to extract the region from the filename, then get
	// (or create) the inventory struct.  We'll append the path of
	// the inventory file to the inventories readQueue.
	for _, p := range paths {
		var inv *Inventory
		if inv = InventoryFromPath(p); inv == nil {
			log.Error("path failed to produce inventory ", p)
			continue
		}
		inv.addPath(p)
	}
}

// InventoryFromPath extracts the region from the filename, then
// returns the corresponding *Inventory.  If the inventory did
// not exist, it will be created.
func InventoryFromPath(path string) *Inventory {
	_, fname := filepath.Split(path)
	if fname == "" {
		log.Error("failed to split path ", path)
		return nil
	}

	// extract the region name and file type from path name
	flen, extlen := len(fname), len(filepath.Ext(path))
	otype := fname[0:3] // type is first 4 chars "vols" or "inst"
	region := fname[4 : flen-extlen]
	if region == "" || otype == "" {
		log.Errorf("expected a region and type got region (%s) and type (%s) ",
			region, otype)
		return nil
	}

	inv := NewInventory(region)
	inv.addPath(path)
	return inv
}

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

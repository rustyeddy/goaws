package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	log "github.com/rustyeddy/logrus"
)

// ======================================================================

// Inventory tracks the responses from requests, the Interface functions
// will be responsible for extracting the desired values from the
// original provider structure.
type Inventory struct {
	Name      string
	Instances map[string]Host // straight from the source
	Volumes   map[string]Disk // from the source
	IV        map[string]string
	VI        map[string]string

	pathQueue  []string                 // inventory files to be scanned
	ec2InsList []ec2.Instance           // from ec2.DescribeInstances
	ec2VolList []ec2.CreateVolumeOutput // from ec2.DescribeVolumes
	*ec2.EC2
	err error
}

// NewInventory willl be created
func NewInventory(name string) Inventory {
	return Inventory{
		Name:      name,
		Instances: make(map[string]Host),
		Volumes:   make(map[string]Disk),
		VI:        make(map[string]string),
		IV:        make(map[string]string),
	}
}

// String print summary of the inventory
func (inv *Inventory) String() string {
	return fmt.Sprintf("%s instances %d - volumes %d ",
		inv.Name, len(inv.Instances), len(inv.Volumes))
}

// Sizes returns the size of Instances and Volumes
func (inv *Inventory) Sizes() (int, int) {
	return len(inv.Instances), len(inv.Volumes)
}

// Save (some of) the inventory
func (inv *Inventory) Save() {

	// Save the entire inventory file!
	jbytes, err := json.Marshal(inv)
	if err != nil {
		log.Fatalf("Failed to json ify my inventory")
	}
	fname := "run/inventory-" + inv.Region + ".json"
	err = ioutil.WriteFile(fname, jbytes, 0644)
	if err != nil {
		log.Fatalf("failed to write the inventory")
	}
}

// Read the inventory from file
func (inv *Inventory) Read(path string) {
	jbytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("failed to marshal JSON")
		return
	}
	err = ioutil.WriteFile(path, jbytes, 0644)
	if err != nil {
		log.Error("failed to write inventory to file")
	}
}

// Index the inventory
func (inv *Inventory) Index() {

	// index the instances
	for _, inst := range inv.Instances {
		iid := *aws.String(*inst.InstanceId)
		vid := *aws.String(*inst.BlockDeviceMappings[0].Ebs.VolumeId)
		inv.IV[iid] = vid
	}

	// read and process volumes
	for _, vol := range inv.Volumes {
		vid := *aws.String(*vol.VolumeId)
		iid := *aws.String(*vol.Attachments[0].InstanceId)
		inv.VI[vid] = iid
	}
}

// indexInstances
func (inv *Inventory) indexInstances(rlist []ec2.RunInstancesOutput) {
	for _, ilist := range rlist {
		for _, inst := range ilist.Instances {
			iid := *inst.InstanceId
			inv.Instances[iid] = HostFromInstance(&inst)
			if inst.BlockDeviceMappings != nil {
				bm0 := inst.BlockDeviceMappings[0]
				ebs := bm0.Ebs
				vid := ebs.VolumeId
				inv.IV[iid] = *vid
			} else {
				inv.IV[iid] = ""
			}
		}
	}
}

// indexVolumes
func (inv *Inventory) indexVolumes(vols []ec2.CreateVolumeOutput) {
	// Index teh volmes and volume to image map
	for _, vol := range vols {
		vid := *vol.VolumeId
		inv.Volumes[vid] = DiskFromVolume(&vol)
		if vol.Attachments != nil {
			a := vol.Attachments[0]
			inv.VI[vid] = *a.InstanceId
		} else {
			inv.VI[vid] = ""
		}
	}
}

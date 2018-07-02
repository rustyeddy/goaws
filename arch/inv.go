package goaws

import (
	"fmt"
	"io/ioutil"

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

	ec2InsList []ec2.Instance           // from ec2.DescribeInstances
	ec2VolList []ec2.CreateVolumeOutput // from ec2.DescribeVolumes
	*ec2.EC2
	err error
}

// IMap is an Inventory map
type IMap map[string]*Inventory

var (
	// Inventories map one *Inventory per region
	inventories IMap // map from region name to Inventory
)

func init() {
	inventories = make(IMap)
}

// NewInventory will create the inventory indexes.  Name is the
// region we are extracting the inventory from.
func NewInventory(name string) *Inventory {
	return &Inventory{
		Name:      name,
		Instances: make(map[string]Host),
		Volumes:   make(map[string]Disk),
		VI:        make(map[string]string),
		IV:        make(map[string]string),
	}
}

// GetInventory will return the specified inventory if it exists, nil if not
func GetInventory(name string) *Inventory {
	var (
		inv *Inventory
		err error
	)

	log.Debugln("~~> GetInventory .. ")
	defer log.Debugln("<~~ Return Inventory .. ")

	log.Debugln("  -- check for a local copy of inventory .. ")
	if inv, ex := inventories[name]; ex {
		log.Debugf("  -- found %s returning %+v", name, inv)
		return inv
	}

	name = "inventory-" + name
	inv = NewInventory(name)
	if inv == nil {
		inventories[name] = nil
		log.Fatalf("Failed to create NewInventory %s ", name)
		return nil
	}

	log.Debug("  -- Find inventory for ", name)
	log.Debug("     -- checking the cache .. ")

	if cache.Exists(name) {
		log.Debug("   -- checking for cached content .. ")
		if err = cache.FetchObject(name, inv); err != nil {
			log.Errorf("  ## failed to fetch content ... ")
			return nil
		}
		return inv
	}
	inv = inv.FetchInventory()

	return inventories[name]
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

	/*

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
	*/
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

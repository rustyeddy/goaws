package goaws

/*

AWS Inventories.  The AWS Inventories are separated by Region and
maintain indexes of Volumes, Instances.  TODO: Snapshots.

*/

import (
	log "github.com/rustyeddy/logrus"
)

// IMap is a map of *Inventory indexed by Region Name.
type IMap map[string]Inventory

var (
	regions     []string // region names
	inventories IMap     // map from region name to Inventory
)

func init() {
	inventories = make(IMap)
}

// Inventories returns the inventory map
func Inventories() IMap {
	return inventories
}

// Exists lets us know if it exists or not
func (im IMap) Exists(name string) bool {
	_, e := im[name]
	return e
}

// Get the specified
func (im IMap) Get(region string) *Inventory {
	if i, e := im[region]; e {
		return &i
	}
	return nil
}

// Set the specified object
func (im IMap) Set(region string, inv *Inventory) {
	im[region] = *inv
}

// Create an inventory and set the index.  Error will be logged and inventory
// will NOT be returned if it already exits.
func (im IMap) Create(region string) (inv *Inventory) {
	var e bool
	if *inv, e = im[region]; e {
		log.Errorf("expected no inventory in region (%s) got (%v) ", region, inv)
		return nil
	}
	return inv
}

// GetOrCreate will get the entry if it exists, if not it will be
// created and indexed before being return to calller.
func (im IMap) GetOrCreate(region string) (inv *Inventory) {
	if i, e := im[region]; !e {
		i = NewInventory(region)
		inv = &i
		im[region] = i
	}
	return inv
}

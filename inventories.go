package main

// Inventories are saved (indexed) by region
var (
	Inventories map[string]Inventory
)

func init() {
	Inventories = make(map[string]Inventory)
}

// GetInventory or create an inventory from the specified region
func GetInventory(name string) *Inventory {
	if i, e := Inventories[name]; e {
		return &i
	}
	// TODO: check that inventories in a region name
	inv := NewInventory(name)
	Inventories[name] = inv
	return &inv
}

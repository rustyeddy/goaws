package goaws

import (
	"github.com/rustyeddy/store"
)

var (
	cache *store.Store
)

// Cache will return the cache
func Cache() *store.Store {
	if cache == nil {
		cache = store.UseStore(config.Basedir)
	}
	return cache
}

// GetInventory from the cache for the given region
func (c *store.Store) GetInventory(region string) (inv *Inventory) {

	inv = c.Get(region)
}

package goaws

import (
	"github.com/rustyeddy/store"
)

var (
	cache *store.Store
	path  string
)

// Cache will return the cache
func Cache() *store.Store {
	if cache == nil {
		cache = store.UseStore(config.Basedir)
	}
	return cache
}

package goaws

import (
	"github.com/rustyeddy/store"
)

var (
	defaultBasedir string = "/srv/www/goaws"
	cache          *store.Store
)

func init() {
	cache = SetCache(defaultBasedir) // set a resonable default
}

// SetCache to whatever dir you want
func SetCache(dir string) *store.Store {
	cache = store.UseStore(dir)
	return cache
}

// Cache returns the goaws cache
func Cache() *store.Store {
	return cache
}

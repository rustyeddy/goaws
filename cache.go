package goaws

import (
	"github.com/rustyeddy/store"
)

var (
	defaultBasedir string = "/srv/www/goaws"
	cache          *store.Store
)

func NewCache(dir string) *store.Store {
	return store.UseStore(dir)
}

// Cache will return the cache
func Cache() *store.Store {
	if cache == nil {
		cache = NewCache(defaultBasedir)
	}
	return cache
}

package goaws

import (
	"github.com/rustyeddy/store"
)

var (
	cache   *store.Store
	basedir string
)

func init() {
	basedir = "/srv/goaws/cache"
	cache = store.UseStore(basedir)
}

// Cache will return the cache
func Cache() *store.Store {
	return cache
}

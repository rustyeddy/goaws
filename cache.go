package goaws

import (
	log "github.com/rustyeddy/logrus"
	"github.com/rustyeddy/store"
)

// Cache object
type Cache struct {
	*store.Store
	*log.Logger // Our private logger
	hits        int
}

var (
	defaultBasedir string = "/srv/www/goaws"
	cache          *Cache
)

func init() {
	cache = NewCache(defaultBasedir) // set a resonable default
}

// SetCache to whatever dir you want
func NewCache(dir string) (cache *Cache) {
	cache = &Cache{store.UseStore(dir), log.New(), 0}
	return cache
}

// Cache returns the goaws cache
func GetCache() (cache *Cache) {
	if cache == nil {
		if cache = NewCache(defaultBasedir); cache == nil {
			log.Errorf("NewCache expected(%s) get ()", defaultBasedir)
			return nil
		}
	}
	return cache
}

// Regions returns the regions from the cache, if we have it
func (c *Cache) Regions() (regions []string) {
	log.Debug("  No copy of Regions in memory: checking the cache... ")

	// Check for a local cache of regions
	if !cache.Exists("regions") {
		log.Infoln("  -- cache entry was not found ... ")
	} else {
		log.Infoln("  ~~> cache object found! Fetching it ...")
		err := cache.FetchObject("regions", &regions)
		if err != nil {
			log.Debugf("  ## error fetching regions %v ..", err)
		}
	}
	return regions
}

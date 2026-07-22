package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache  map[string]cacheEntry
	secure *sync.Mutex // make safe for concurrent use
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte // represents the raw data being cached
}

// New cache with configurable interval
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache:  make(map[string]cacheEntry),
		secure: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.secure.Lock()         // lock the mutex while adding to cache
	defer c.secure.Unlock() // set to unlock when finished

	c.cache[key] = cacheEntry{ //
		createdAt: time.Now().UTC(), // set time of creation in cache
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.secure.Lock()
	defer c.secure.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(elapsed time.Duration) {
	ticker := time.NewTicker(elapsed) // start a time counter
	for range ticker.C {
		timeNow := time.Now().UTC()
		c.reap(timeNow, elapsed)
	}
}

func (c *Cache) reap(now time.Time, passed time.Duration) {
	c.secure.Lock() // lock the mutex while harvesting from cache
	defer c.secure.Unlock()

	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-passed)) {
			delete(c.cache, key)
		}
	}
}

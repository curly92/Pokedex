package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry    map[string]cacheEntry
	mut      sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	newEntry := map[string]cacheEntry{}
	cache := &Cache{
		newEntry,
		sync.Mutex{},
		duration,
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	c.entry[key] = cacheEntry{
		time.Now(),
		val,
	}
	c.mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	entry, ok := c.entry[key]
	if !ok {
		c.mut.Unlock()
		return nil, false
	}
	val := entry.val
	c.mut.Unlock()
	return val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-c.interval)

		c.mut.Lock()
		for k, e := range c.entry {
			if e.createdAt.Before(cutoff) {
				delete(c.entry, k)
			}
		}
		c.mut.Unlock()
	}
}

package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	*sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		map[string]cacheEntry{},
		&sync.Mutex{},
		interval,
	}
	return cache
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		time.Now(),
		val,
	}
	c.Lock()
	defer c.Unlock()
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.Lock()
	defer c.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		t := <-ticker.C
		c.Lock()
		for key, entry := range c.cache {
			if entry.createdAt.Before(t) {
				delete(c.cache, key)
			}
		}
		c.Unlock()
	}
}
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	lock *sync.Mutex
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
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
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
		for key, entry := range c.cache {
			if entry.createdAt.Before(t) {
				delete(c.cache, key)
			}
		}
	}
}
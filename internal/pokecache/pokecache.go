package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu *sync.Mutex
	cache map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		&sync.Mutex{},
		map[string]cacheEntry{},
		interval,
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		time.Now(),
		val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = entry
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
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
		c.mu.Lock()
		for key, entry := range c.cache {
			if entry.createdAt.Before(t) {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
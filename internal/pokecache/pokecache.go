package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	lock *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}
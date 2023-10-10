package pokecache

import (
	"sync"
	"time"
)

// Main struct for the cache
type Cache struct {
	cache  map[string]cacheEntry
	mu     *sync.Mutex
}

// Represents a single cache entry
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

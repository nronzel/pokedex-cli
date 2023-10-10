package pokecache

import (
	"sync"
	"time"
)

// Main struct for the cache
type Cache struct {
	items  map[string]cacheEntry
	mu     sync.Mutex
	ticker *time.Ticker
}

// Represents a single cache entry
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

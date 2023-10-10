package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		items:  make(map[string]cacheEntry),
		ticker: time.NewTicker(interval),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	for {
		<-c.ticker.C
		c.mu.Lock()
		for key, entry := range c.items {
			if time.Since(entry.createdAt) > interval {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.items[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

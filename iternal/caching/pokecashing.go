package cashing

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	caches map[string]cacheEntry
	mu     *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		caches: make(map[string]cacheEntry),
		mu:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.caches[key] = entry
}

func (c *Cache) Get(key string) (res []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exists := c.caches[key]; exists {
		return c.caches[key].val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		defer c.mu.Lock()
		for key, element := range c.caches {
			if time.Since(element.createdAt) > interval {
				delete(c.caches, key)
			}
		}
	}

}

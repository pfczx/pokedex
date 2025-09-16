package cashing

import (
	"sync"
	"time"
)

type casheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cashes map[string]casheEntry
	mu     *sync.Mutex
}

func NewCashe(interval time.Duration) *Cache {
	c := &Cache{
		cashes: make(map[string]casheEntry),
		mu:     &sync.Mutex{},
	}
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := casheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cashes[key] = entry
}



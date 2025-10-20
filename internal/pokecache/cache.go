package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	entries map[string]CacheEntry
	ttl     time.Duration
	mux     sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entries[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	if time.Since(entry.CreatedAt) > c.ttl {
		delete(c.entries, key)
		return nil, false
	}
	return entry.Val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mux.Lock()
		for key, entry := range c.entries {
			if now.Sub(entry.CreatedAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}

func NewCache(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]CacheEntry),
		ttl:     ttl,
		mux:     sync.RWMutex{},
	}
	go c.reapLoop()

	return c
}

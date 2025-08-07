package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}
	c.mutex.Unlock()
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mutex.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}
}
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: map[string]cacheEntry{},
		mutex: sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

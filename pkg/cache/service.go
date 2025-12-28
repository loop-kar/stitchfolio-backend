package cache

import (
	"sync"
	"time"
)

func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, 24*time.Hour)
}

func (c *Cache) SetWithTTL(key string, value interface{}, expiration time.Duration) {
	item := CacheStructure{
		Value: value,
	}
	if expiration > 0 {
		item.Expiration = time.Now().Add(expiration)
	}
	c.store.Store(key, item)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if value, ok := c.store.Load(key); ok {
		item := value.(CacheStructure)

		// perfrom check
		if !item.Expiration.IsZero() && time.Now().After(item.Expiration) {
			c.store.Delete(key)
			return nil, false
		}
		return item.Value, true
	}
	return nil, false
}

func (c *Cache) Delete(key string) {
	c.store.Delete(key)
}

func (c *Cache) Clear() {
	c.store = sync.Map{}
}

// GetOrSet retrieves a value from the cache, or sets it if it doesn't exist
func (c *Cache) GetOrSet(key string, value interface{}, expiration time.Duration) interface{} {
	if val, ok := c.Get(key); ok {
		return val
	}
	c.SetWithTTL(key, value, expiration)
	return value
}

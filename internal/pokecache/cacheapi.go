package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{}
	ticker := time.NewTicker(interval)

	tickerChan := make(chan bool)

	go func() {
		for {
			select {
			case <-tickerChan:
				return
			case <-ticker.C:
				newCache.reapLoop(interval)
			}
		}
	}()

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	if c.cacheMap == nil {
		c.cacheMap = make(map[string]cacheEntry)
	}
	if _, ok := c.cacheMap[key]; !ok {
		c.cacheMap[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for key, entry := range c.cacheMap {
		if time.Since(entry.createdAt) > interval {
			delete(c.cacheMap, key)
		}
	}
}

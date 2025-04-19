package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex    *sync.Mutex
	entries  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mutex: &sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}
	
	go cache.reapLoop(interval)
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry := cacheEntry{
		createdAt: time.Now().UTC(),
		value:     val,
	}

	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry, exists := cache.entries[key]
	return entry.value, exists
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	for key, entry:= range cache.entries {
		if entry.createdAt.Before(now.Add(-last)) {
			delete(cache.entries, key)
		}
	}
}

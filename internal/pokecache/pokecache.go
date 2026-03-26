package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval int) (*Cache, error) {
	cache := Cache{}

	reapLoop(cache, interval)
	return cache, nil
}

func Add(cache *Cache) error {
	return nil
}

func Get(cache *Cache) ([]byte, bool) {
	return nil, false
}

func reapLoop(cache *Cache, interval time.Duration) {

}

package store

import (
	"fmt"
	"sync"

	Ca "github.com/theweird-kid/cache-go/internals/cache"
)

type Store struct {
	data  map[int]string
	cache Ca.Cacher
	mu    sync.RWMutex
}

func NewStore(c Ca.Cacher) *Store {
	data := map[int]string{
		1: "Hi this is Gaurav",
		2: "this is his Pub/sub app",
	}

	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Get(key int) (string, error) {
	s.mu.RLock()
	val, ok := s.cache.Get(key)
	s.mu.RUnlock()
	if ok {
		//burst the cache
		s.mu.Lock()
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		s.mu.Unlock()
		fmt.Println("returning the value from cache..")
		return val, nil
	}

	s.mu.RLock()
	val, ok = s.data[key]
	s.mu.RUnlock()
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	// set the new value into cache
	s.mu.Lock()
	if err := s.cache.Set(key, val); err != nil {
		s.mu.Unlock()
		fmt.Println(err)
		return "", err
	}
	s.mu.Unlock()

	fmt.Println("returning the value from internal storage..")
	return val, nil
}

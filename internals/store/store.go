package store

import (
	"fmt"
	"sync"

	Ca "github.com/theweird-kid/cache-go/internals/cache"
)

type Store struct {
	data  map[int]string
	cache Ca.Cacher
	mu    sync.Mutex
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
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.cache.Get(key)
	if ok {
		// burst the cache
		s.cache.Remove(key)
		fmt.Println("returning data from cache")
		return val, nil
	}

	// If not found in cache, get from data map
	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}

	fmt.Println("returning data from internal storage")
	s.cache.Set(key, val)
	return val, nil
}

func (s *Store) Set(key int, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Set the value in the data map
	fmt.Println("Internal storage updated")
	s.data[key] = value

	// Optionally, set the value in the cache
	fmt.Println("cache updated")
	s.cache.Set(key, value)
	return nil
}

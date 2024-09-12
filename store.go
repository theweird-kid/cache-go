package main

import "fmt"

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
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
	val, ok := s.cache.Get(key)
	if ok {
		//burst the cache
		if err := s.cache.Remove(key); err != nil {
			fmt.Println(err)
		}
		fmt.Println("returning the value from cache..")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("returning the value from internal storage..")
	return val, nil
}

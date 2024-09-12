package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

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
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}
	return val, nil
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println(client)

	s := NewStore(NopCache{})
	val, err := s.Get(1)
	if err != nil {
		log.Fatal()
	}

	fmt.Println(val)
}

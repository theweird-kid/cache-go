package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

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

package main

import (
	"context"
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

	ctx := context.Background()

	for i := 0; i < 100; i++ {
		if err := client.Publish(ctx, "coordsTopic", i).Err(); err != nil {
			log.Fatal(err)
		}
	}

}

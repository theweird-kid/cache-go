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

	// Subscriber
	ctx := context.Background()
	sub := client.Subscribe(ctx, "coordsTopic")
	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", msg)
	}
}

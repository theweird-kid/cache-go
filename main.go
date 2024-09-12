package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	cache := store.NewInMemoryCache() // Assuming you have an in-memory cache implementation
	store := store.NewStore(cache)
	pb.RegisterStoreServiceServer(s, &server{store: store})

	fmt.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/theweird-kid/cache-go/internals/cache"
	"github.com/theweird-kid/cache-go/internals/store"
	pb "github.com/theweird-kid/cache-go/proto/serverpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Entry Point")
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR environment variable not set")
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	fmt.Println(client)

	cache := cache.NewRedisCache(client, time.Second*5)
	myStore := store.NewStore(cache)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStoreServiceServer(s, NewgRPCServer(myStore))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

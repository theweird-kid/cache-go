package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/theweird-kid/cache-go/internals/cache"
	"github.com/theweird-kid/cache-go/internals/store"
	pb "github.com/theweird-kid/cache-go/proto/serverpb"
	"google.golang.org/grpc"
)

type server struct {
	store *store.Store
	pb.UnimplementedStoreServiceServer
}

func NewgRPCServer(store *store.Store) *server {
	return &server{
		store: store,
	}
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// Implement your logic here
	key := req.Key
	data, err := s.store.Get(int(key))
	if err != nil {
		return &pb.GetResponse{Value: "example_value", Source: "cache"}, err
	}
	return &pb.GetResponse{Value: data}, nil
}

func (s *server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	// Implement your logic here
	key := req.Key
	val := req.Value
	err := s.store.Set(int(key), val)
	fmt.Println(err)
	if err != nil {
		return &pb.SetResponse{Success: false}, err
	}
	return &pb.SetResponse{Success: true}, nil
}

func main() {
	fmt.Println("Entry Point")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
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

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/theweird-kid/cache-go/internals/store"
	pb "github.com/theweird-kid/cache-go/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStoreServiceServer
	store *store.Store
}

func (s *server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	key := int(req.GetKey())
	val, err := s.store.Get(key)
	if err != nil {
		return nil, err
	}
	source := "internal storage"
	if _, ok := s.store.Cache.Get(key); ok {
		source = "cache"
	}
	return &pb.GetResponse{Value: val, Source: source}, nil
}

func (s *server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	key := int(req.GetKey())
	value := req.GetValue()
	err := s.store.Set(key, value)
	if err != nil {
		return &pb.SetResponse{Success: false}, err
	}
	return &pb.SetResponse{Success: true}, nil
}

func main() {
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

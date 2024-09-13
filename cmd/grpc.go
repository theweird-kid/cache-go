package main

import (
	"context"
	"fmt"

	"github.com/theweird-kid/cache-go/internals/store"
	pb "github.com/theweird-kid/cache-go/proto/serverpb"
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

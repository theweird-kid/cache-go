package main

import (
	"context"

	"github.com/theweird-kid/go-cache/internals/store"
	pb "github.com/theweird-kid/go-cache/proto"
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

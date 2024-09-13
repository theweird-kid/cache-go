package main

import (
	"context"
	"log"
	"time"

	pb "github.com/theweird-kid/cache-go/proto/serverpb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStoreServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call Set method

	setResp, err := c.Set(ctx, &pb.SetRequest{Key: 1, Value: "example_value"})
	if err != nil {
		log.Fatalf("could not set value: %v", err)
	}
	log.Printf("Set Response: %v", setResp.Success)

	// Call Get method
	getResp, err := c.Get(ctx, &pb.GetRequest{Key: 1})
	if err != nil {
		log.Fatalf("could not get value: %v", err)
	}
	log.Printf("Get Response: %s", getResp.Value)
}

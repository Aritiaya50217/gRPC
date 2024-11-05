package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Aritiaya50217/gRPC/grpc-hello-world/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	creds := insecure.NewCredentials()

	// Create a new gRPC client connection (ต่อไปยังที่เดียวกับ gRPC server)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to create client : %v ", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := "Test"
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet : %v ", err)
	}
	log.Printf("Greeting : %s", r.GetMessage())

}

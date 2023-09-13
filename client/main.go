package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ignite91/simple-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAllUsers(ctx, &pb.GetAllUsersRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("GetAllUsers: %s", r)
}

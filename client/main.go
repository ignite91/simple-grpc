package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ignite91/simple-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	t := time.Now()
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServicesClient(conn)
	ctx := context.Background()
	/* 	r, err := c.GetAllUsers(ctx, &pb.GetAllUsersRequest{})
	   	if err != nil {
	   		log.Fatalf("could not getAllUsers: %v", err)
	   	}
	   	fmt.Println("MSG: ", r.User) */
	in := &pb.SaveUsersRequest{
		User: []*pb.User{{
			Id:       1,
			Name:     "name",
			Lastname: "lastName",
			Age:      99,
			Active:   true,
			Money:    33999.22,
			Saveat:   "saveAt",
		},
		}}

	for i := 0; i < 5000; i++ {
		_, err := c.SaveUsers(ctx, in)
		if err != nil {
			log.Fatalf("could not saveUsers: %v", err)
		}
	}
	fmt.Println("Elapsed: ", time.Since(t).Seconds())
}

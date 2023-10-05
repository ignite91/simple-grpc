package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	pb "github.com/ignite91/simple-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	t := time.Now()
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewServicesClient(conn)
	ctx := context.Background()
	/* 	for i := 0; i < 20000; i++ {
		_, err := c.GetOne(ctx, &pb.GetOneRequest{})
		if err != nil {
			log.Fatalf("could not getAllUsers: %v", err)
		}
	} */
	in := &pb.GetAllRequest{}
	_, err = c.GetAll(ctx, in, grpc.MaxCallRecvMsgSize(math.MaxInt32))
	if err != nil {
		log.Fatalf("could not GetAll: %v", err)
	}
	fmt.Println("Elapsed from client:", time.Since(t).Seconds())
}

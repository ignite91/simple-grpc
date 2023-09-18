package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/ignite91/simple-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedServicesServer
}

func (s *server) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersReply, error) {
	data, err := os.ReadFile("users.proto")
	if err != nil {
		panic(err)
	}
	out := &pb.GetAllUsersReply{}
	err = proto.Unmarshal(data, out)
	if err != nil {
		panic(err)
	}
	return out, nil
}

// COMO USER SEPARADO
/* func (s *server) SaveUsers(ctx context.Context, in *pb.SaveUsersRequest) (*pb.SaveUsersUsersReply, error) {
	f, err := os.OpenFile("users.proto", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	msg, err := proto.Marshal(in)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(msg)
	if err != nil {
		panic(err)
	}
	return &pb.SaveUsersUsersReply{}, nil
} */
// COMO ARRAY
func (s *server) SaveUsers(ctx context.Context, in *pb.SaveUsersRequest) (*pb.SaveUsersUsersReply, error) {
	data, err := os.ReadFile("users.proto")
	if err != nil {
		panic(err)
	}
	users := &pb.SaveUsersRequest{}
	err = proto.Unmarshal(data, users)
	if err != nil {
		panic(err)
	}
	proto.Merge(users, in)
	usersb, err := proto.Marshal(users)
	if err != nil {
		panic(err)
	}
	os.WriteFile("users.proto", usersb, 0644)
	if err != nil {
		panic(err)
	}
	return &pb.SaveUsersUsersReply{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServicesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

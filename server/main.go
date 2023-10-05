package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ignite91/simple-grpc/proto"
	"google.golang.org/grpc"
)

var model = &pb.Model{
	Campo1:  99999999.22,
	Campo2:  99999999.22,
	Campo3:  99999999.22,
	Campo4:  99999999.22,
	Campo5:  99999999.22,
	Campo6:  99999999.22,
	Campo7:  99999999.22,
	Campo8:  99999999.22,
	Campo9:  99999999.22,
	Campo10: 99999999.22,
	Campo11: 99999999.22,
	Campo12: 99999999.22,
	Campo13: 99999999.22,
	Campo14: 99999999.22,
	Campo15: 99999999.22,
	Campo16: 99999999.22,
	Campo17: 99999999.22,
	Campo18: 99999999.22,
	Campo19: 99999999.22,
	Campo20: 99999999.22,
	Campo21: 99999999.22,
	Campo22: 99999999.22,
	Campo23: 99999999.22,
	Campo24: 99999999.22,
	Campo25: 99999999.22,
	Campo26: 99999999.22,
	Campo27: 99999999.22,
	Campo28: 99999999.22,
	Campo29: 99999999.22,
	Campo30: 99999999.22,
	Campo31: 99999999.22,
	Campo32: 99999999.22,
	Campo33: 99999999.22,
	Campo34: 99999999.22,
	Campo35: 99999999.22,
	Campo36: 99999999.22,
	Campo37: 99999999.22,
	Campo38: 99999999.22,
	Campo39: 99999999.22,
	Campo40: 99999999.22,
	Campo41: 99999999.22,
	Campo42: 99999999.22,
	Campo43: 99999999.22,
	Campo44: 99999999.22,
	Campo45: 99999999.22,
	Campo46: 99999999.22,
	Campo47: 99999999.22,
	Campo48: 99999999.22,
	Campo49: 99999999.22,
	Campo50: 99999999.22,
	Campo51: 99999999.22,
	Campo52: 99999999.22,
	Campo53: 99999999.22,
	Campo54: 99999999.22,
	Campo55: 99999999.22,
	Campo56: 99999999.22,
	Campo57: 99999999.22,
	Campo58: 99999999.22,
	Campo59: 99999999.22,
	Campo60: 99999999.22,
}

type server struct {
	pb.UnimplementedServicesServer
}

func (s *server) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllReply, error) {
	var models []*pb.Model
	for i := 0; i < 20000; i++ {
		models = append(models, model)
	}
	out := &pb.GetAllReply{
		Model: models,
	}
	return out, nil
}
func (s *server) GetOne(ctx context.Context, in *pb.GetOneRequest) (*pb.GetOneReply, error) {
	out := &pb.GetOneReply{
		Model: model,
	}
	return out, nil
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

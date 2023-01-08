package rpc01

import (
	"context"
	"flag"
	"fmt"
	pb "go_primer/grpc/gen-go/grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserIconServer
}

func (s *server) GetUserIcon(context.Context, *pb.UserIconRequest) (*pb.UserIconResponse, error) {
	return &pb.UserIconResponse{}, nil
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserIconServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

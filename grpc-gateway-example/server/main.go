package main

import (
	"log"
	"net"

	pb "github.com/xingcuntian/go_test/grpc-gateway-example/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

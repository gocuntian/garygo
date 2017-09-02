package main

import (
	"fmt"
	"log"
	"net"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	proto "github.com/xingcuntian/go_test/go-microservice/grpc/proto"
)

type kittenServer struct{}

func (k *kittenServer) Hello(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	response := new(proto.Response)
	response.Msg = fmt.Sprintf("Hello %v", request.Name)
	return response, nil
}

func main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterKittensServer(grpcServer, &kittenServer{})
	grpcServer.Serve(l)
}

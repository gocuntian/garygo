package main

import (
	"log"

	micro "github.com/micro/go-micro"
	proto "github.com/xingcuntian/go_test/go-micro-example/stream/server/proto"
	"golang.org/x/net/context"
)

type Streamer struct{}

func (e *Streamer) ServerStream(ctx context.Context, req *proto.Request, stream proto.Streamer_ServerStreamStream) error {
	log.Printf("Got msg %v", req.Count)
	for i := 0; i < int(req.Count); i++ {
		if err := stream.Send(&proto.Response{Count: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

func (e *Streamer) Stream(ctx context.Context, stream proto.Streamer_StreamStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got msg %v", req.Count)
		if err := stream.Send(&proto.Response{Count: req.Count}); err != nil {
			return err
		}
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.stream"),
	)
	service.Init()
	proto.RegisterStreamerHandler(service.Server(), new(Streamer))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

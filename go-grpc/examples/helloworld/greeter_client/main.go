package main

import (
	"log"
	// "os"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/xingcuntian/go_test/go-grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
)

func main() {
	
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	var id int64 = 1
	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
	//log.Printf("Greeting: %s", r.Message)
}

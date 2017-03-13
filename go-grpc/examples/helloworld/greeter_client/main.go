package main

import (
	"log"
	"io"
	// "os"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/xingcuntian/go_test/go-grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
)

func getCustomers(client pb.GreeterClient, filter *pb.CustomerFilter){
	stream, err :=client.GetCustomers(context.Background(),filter);
	if err!=nil{
		log.Fatalf("Error on get customers:%v",err)
	}
	for {
		customer, err :=stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v",customer)
	}
}

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

   
	// Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: ""}
	getCustomers(c, filter)
}

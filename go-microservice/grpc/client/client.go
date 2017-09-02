package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/xingcuntian/go_test/go-microservice/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	client := proto.NewKittensClient(con)
	response, err := client.Hello(context.Background(), &proto.Request{Name: "xingcuntian"})
	if err != nil {
		log.Fatal("Error calling service: ", err)
	}

	fmt.Println(response.Msg)
}

package main

import (
    pb "github.com/xingcuntian/go_test/go-grpc-example/example/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

const (
    Address = "127.0.0.1:50052"
)

func main(){
    conn,err:=grpc.Dial(Address,grpc.WithInsecure())
    if err!=nil{
       grpclog.Fatalf("failed to conn: %v",err)
    }
    defer conn.Close()
    //初始化客户端
    c:=pb.NewHelloClient(conn)
    reqBody:=new(pb.HelloRequest)
    reqBody.Name = "this is gRPC test"
    r,err:=c.SayHello(context.Background(),reqBody)
    if err!=nil{
        grpclog.Fatalf("failed to reqBody: %v",err)
    }
    grpclog.Println(r.Messge)
}
package main

import (
    pb "github.com/xingcuntian/go_test/go-grpc-example/example/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials" // 引入grpc认证包
    "google.golang.org/grpc/grpclog"
)

const (
    Address = "127.0.0.1:50052"
)

func main(){
    //TLS 连接
    creds,err:=credentials.NewClientTLSFromFile("../../keys/server.pem", "server name")
    if err!=nil{
        grpclog.Fatalf("Failed to create TLS credentials %v", err)
    }

    conn,err:=grpc.Dial(Address,grpc.WithTransportCredentials(creds))
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
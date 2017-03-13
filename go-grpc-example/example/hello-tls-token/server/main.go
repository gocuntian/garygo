package main

import (
    "fmt"
    "net"
    pb "github.com/xingcuntian/go_test/go-grpc-example/example/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/metadata"
)

const (
    //Address gRPC服务地址
    Address = "127.0.0.1:50052"
)
//定义helloService并实现约定的接口
type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    //解析metadata中的信息并验证
    md, ok := metadata.FromContext(ctx)
    if !ok{
        return nil, grpc.Errorf(codes.Unauthenticated,"无token认证信息")
    }

    var (
        appid string
        appkey string
    )

     if val, ok := md["appid"]; ok {
        appid = val[0]
    }

    if val, ok := md["appkey"]; ok {
        appkey = val[0]
    }

    if appid != "101010" || appkey != "i am key" {
        return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
    }

    resp := new(pb.HelloReply)

    resp.Messge = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)

    return resp, nil
}

func main(){
    listen,err:=net.Listen("tcp",Address)
    if err!=nil{
        grpclog.Fatalf("failed to listen: %v",err)
    }
    //TLS认证
    creds,err:=credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
    if err!=nil{
        grpclog.Fatalf("Failed to generate credentials %v", err)
    }


    //实例化gRPC Server
    s:=grpc.NewServer(grpc.Creds(creds))



    //注册HelloService
    pb.RegisterHelloServer(s,HelloService)   

    grpclog.Println("Listen on " + Address + " with TLS + Token")
    s.Serve(listen)
}
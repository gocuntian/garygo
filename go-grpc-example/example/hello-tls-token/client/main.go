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
    //是否开启TLS认证
    OpenTLS = true
)

//自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error){
    return map[string]string{
        "appid":"1010100",
        "appkey":"i am key",
    },nil
}

func (c customCredential)RequireTransportSecurity()bool{
    if OpenTLS {
        return true
    }
    return false
}

func main(){
    var err error
    var opts []grpc.DialOption
    if OpenTLS{
        //TLS 连接
        creds,err:=credentials.NewClientTLSFromFile("../../keys/server.pem", "server name")
        if err!=nil{
            grpclog.Fatalf("Failed to create TLS credentials %v", err)
        }
        opts = append(opts,grpc.WithTransportCredentials(creds))
    }else{
        opts = append(opts,grpc.WithInsecure())
    }

    //使用自定义认证
    opts = append(opts,grpc.WithPerRPCCredentials(new(customCredential)))    

    conn,err:=grpc.Dial(Address,opts...)
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
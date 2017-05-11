package main

import (
    "github.com/xingcuntian/go_test/go-example/http/rpc/rpcexample"
    "log"
    "net/rpc"
)

// func DialHTTP(network, address string) (*Client, error)
// DialHTTP在指定的网络和地址与在默认HTTP RPC路径监听的HTTP RPC服务端连接。

// func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
// Call调用指定的方法，等待调用返回，将结果写入reply，然后返回执行的错误状态。
func main(){
    client, err := rpc.DialHTTP("tcp",":1234")
    if err != nil {
        log.Fatalf("Error in dialing. %s", err)
    }
    args := &rpcexample.Args{
        A: 2,
        B: 3,
    }

    var result rpcexample.Result
    err = client.Call("Arith.Multiply",args,&result)
    if err != nil {
         log.Fatalf("error in Arith", err)
    }
    log.Printf("%d * %d = %d\n",args.A,args.B,result)
}
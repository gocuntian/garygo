package main

import (
    "github.com/xingcuntian/go_test/go-example/http/rpc/rpcexample"
    "log"
    "net"
    "net/http"
    "net/rpc"
)

// func Fatalf(format string, v ...interface{})
// Fatalf等价于{Printf(v...); os.Exit(1)}

// func Register(rcvr interface{}) error
// Register在DefaultServer注册并公布rcvr的方法。
// func HandleHTTP()
// HandleHTTP函数注册DefaultServer的RPC信息HTTP处理器对应到DefaultRPCPath，和DefaultServer的debug处理器对应到DefaultDebugPath。
// HandleHTTP函数会注册到http.DefaultServeMux。之后，仍需要调用http.Serve()，一般会另开线程："go http.Serve(l, nil)"

// func Listen(net, laddr string) (Listener, error)
// 返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络

// func Serve(l net.Listener, handler Handler) error
// Serve会接手监听器l收到的每一个连接，并为每一个连接创建一个新的服务go程。该go程会读取请求，然后调用handler回复请求。handler参数一般会设为nil，此时会使用DefaultServeMux。
func main(){
    arith := new(rpcexample.Arith)
    err := rpc.Register(arith)
    if err != nil {
        log.Fatalf("Format of service Arith isn't correct. %s", err)
    }
    rpc.HandleHTTP()

    l, e :=net.Listen("tcp",":1234")
    if e != nil {
         log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
    }
    log.Println("Serving RPC handler")
    err = http.Serve(l,nil)
    if err != nil {
        log.Fatalf("Error serving : %s",err)
    }

}
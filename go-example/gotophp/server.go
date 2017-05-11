package main

import (
    "fmt"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
)

type Serve struct{}

func (s *Serve) Echo(a *int, r *int) error {
   *r = 1
   return nil
}

func main(){
    tcpAddr, err := net.ResolveTCPAddr("tcp",":9002")
    if err != nil {
        fmt.Println(err)
        return    
    }

    l, err := net.ListenTCP("tcp",tcpAddr)
    if err != nil {
        fmt.Println(err)
        return
    }

    serve := &Serve{}
    rpc.Register(serve)
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go jsonrpc.ServeConn(conn)
    }
}
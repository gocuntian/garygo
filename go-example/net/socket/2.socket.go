package main

import (
    "io"
    "log"
    "net"
)

func main(){
    listener, _ := net.Listen("tcp",":2000")
    defer listener.Close()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
            continue
        }

        go func(c net.Conn){
            io.Copy(c,c)
            c.Close()
        }(conn)
    }
}
package main

import (
    "net"
    "io"
)
//UNIX套接字的API
func main(){
    unixAddr, _ := net.ResolveUnixAddr("unix","/tmp/echo.sock")

    unixListener, _ := net.ListenUnix("unix",unixAddr)

    for {
        conn, err := unixListener.Accept()
        if err != nil {
            continue
        }
        go func(c net.Conn){
            io.Copy(c,c)
            c.Close()
        }(conn)
    }
}
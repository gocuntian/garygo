package main

import (
    "fmt"
    "net"
    "os"
)

func main(){
    var buf [512]byte
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    udpAddr, _ := net.ResolveUDPAddr("udp",service)
    conn, _ := net.DialUDP("udp",nil,udpAddr)
    defer conn.Close()
    _ , _ = conn.Write([]byte("Hello Server!"))
    n, _ := conn.Read(buf[0:])
    fmt.Println("Reply from server ",conn.RemoteAddr().String(),string(buf[0:n]))
    os.Exit(0)
}
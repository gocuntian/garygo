package main

import (
    "fmt"
    "net"
)

func main(){
    ip := net.ParseIP("224.0.0.250")
    srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port : 0}
    dstAddr := &net.UDPAddr{IP: ip,Port:9981}
    conn, err := net.DialUDP("udp",srcAddr,dstAddr)
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()
    conn.Write([]byte("hello"))
    fmt.Printf("<%s>\n",conn.RemoteAddr())
}
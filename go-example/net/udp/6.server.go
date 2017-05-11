package main

import (
    "fmt"
    "net"
)
//广播

func main(){
    listener, _ := net.ListenUDP("udp",&net.UDPAddr{IP:net.IPv4zero,Port:9981})
    fmt.Printf("Local: <%s> \n",listener.LocalAddr().String())

    data := make([]byte,1024)
    for {
        n, remoteAddr, _ := listener.ReadFromUDP(data)
        fmt.Printf("<%s> %s\n",data[:n],remoteAddr)
        _,_ = listener.WriteToUDP([]byte("hello world!"),remoteAddr)
    }
}
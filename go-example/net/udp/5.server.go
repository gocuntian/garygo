package main

import (
    "fmt"
    "net"
)
//标准库多播编程

func main(){
    //如果第二参数为nil,它会使用系统指定多播接口，但是不推荐这样使用
    addr, err := net.ResolveUDPAddr("udp","224.0.0.250:9981")
    if err != nil {
        fmt.Println(err)
    }

    listener, err := net.ListenMulticastUDP("udp",nil,addr)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
    data := make([]byte,1024)
    for {
        n, remoteAddr, err := listener.ReadFromUDP(data)
        if err != nil {
            fmt.Printf("error during read: %s",err)
        }
        fmt.Printf("<%s> %s\n",remoteAddr,data[:n])
    }
}

// func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
// ListenMulticastUDP接收目的地是ifi接口上的组地址gaddr的UDP数据包。它指定了使用的接口，如果ifi是nil，将使用默认接口。
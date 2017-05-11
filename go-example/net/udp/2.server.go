package main

import (
    "fmt"
    "net"
)

func main(){
    listener, err := net.ListenUDP("udp",&net.UDPAddr{IP: net.ParseIP("127.0.0.1"),Port:9981})
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
        fmt.Printf("<%s> %s \n",remoteAddr,data[:n])
        _, err = listener.WriteToUDP([]byte("hello world!"),remoteAddr)
        if err != nil {
            fmt.Printf(err.Error())
        }
    }

}

// func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
// ReadFromUDP从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。
// ReadFromUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。

// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
// WriteToUDP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。
// WriteToUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。
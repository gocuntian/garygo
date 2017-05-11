package main

import (
    "fmt"
    "net"
    "os"
)
//客户端代码有所不同，它不是通过DialUDP “连接” 广播地址，而是通过ListenUDP创建一个unconnected的 *UDPConn,
//然后通过WriteToUDP发送数据报，这和你脑海中的客户端不太一致
func main(){
    ip := net.ParseIP("172.20.8.32")
    srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port:0}
    dstAddr := &net.UDPAddr{IP: ip, Port:9981}

    conn, _ := net.ListenUDP("udp",srcAddr)
    _, _ = conn.WriteToUDP([]byte("who"),dstAddr)

    data := make([]byte,1024)
    n, addr, _ :=conn.ReadFrom(data)
    fmt.Printf("read %s from <%s>\n",data[:n],addr.String())
    b := make([]byte,1)
    os.Stdin.Read(b)
}
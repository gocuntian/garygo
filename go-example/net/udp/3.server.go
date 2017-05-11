package main

import (
    "fmt"
    "net"
    "os"
    "time"
)
// 两个服务器通信的例子，互为客户端和服务器，在发送数据报的时候，
// 将发送的一方称之为源地址，发送的目的地一方称之为目标地址
func read(conn *net.UDPConn){
    for{
        data := make([]byte,1024)
        n, remoteAddr, err := conn.ReadFromUDP(data)
        if err != nil {
            fmt.Printf("error during read: %s\n",err)
        }
        fmt.Printf("receive %s from <%s>\n",data[:n],remoteAddr)
    }
}

func main(){
    addr1 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port:9001}
    addr2 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port:9002}

    go func(){
        listener1, err := net.ListenUDP("udp",addr1)
        if err != nil {
            fmt.Println(err)
            return
        }
        go read(listener1)

        time.Sleep(5 * time.Second)
        listener1.WriteToUDP([]byte("ping to #2: "+ addr2.String()),addr2)
    }()

    go func(){
        listener2, err := net.ListenUDP("udp",addr2)
        if err != nil {
            fmt.Println(err)
            return
        }
        go read(listener2)
        time.Sleep(5 * time.Second)

        listener2.WriteToUDP([]byte("ping to #1 :" + addr1.String()),addr1)
    }()

    b := make([]byte,1)
    os.Stdin.Read(b)
}


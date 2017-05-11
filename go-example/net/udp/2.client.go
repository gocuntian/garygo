package main

import (
    "fmt"
    "net"
    "os"
)

func main(){
    sip := net.ParseIP("127.0.0.1")
    srcAddr := &net.UDPAddr{IP: net.IPv4zero,Port:0}
    dstAddr := &net.UDPAddr{IP: sip, Port:9981}

    conn, err := net.DialUDP("udp",srcAddr,dstAddr)
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()


    //  b:=make([]byte,1)
    //  os.Stdin.Read(b)

    var message string;

    fmt.Scanln(&message)
    conn.Write([]byte(message+"\n"))
    recData := make([]byte,1024)
    n, _ := conn.Read(recData)

    fmt.Printf("<%s> %s\n",conn.RemoteAddr(),string(recData[:n]))


}
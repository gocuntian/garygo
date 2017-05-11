package main

import (
    "fmt"
    "net"
    "golang.org/x/net/ipv4"
)

// 通用多播编程

// Go标准库也支持多播，但是我们首先我们看通用的多播是如何实现的，它使用golang.org/x/net/ipv4或者golang.org/x/net/ipv6进行控制。

// 首先找到要进行多播所使用的网卡,然后监听本机合适的地址和服务端口。
// 将这个应用加入到多播组中，它就可以从组中监听包信息，当然你还可以对包传输进行更多的控制设置。
// 应用收到包后还可以检查包是否来自这个组的包。

func main(){
    //1. 得到一个interface
    en4, err := net.InterfaceByName("en4")
    if err != nil {
        fmt.Println(err)
    }
    group := net.IPv4(244,0,0,250)

    //2. bind一个本地地址
    c, err := net.ListenPacket("udp4","0.0.0.0:1024")
    if err != nil {
        fmt.Println(err)
    }
    defer c.Close()

    p := ipv4.NewPackConn(c)
    if err := p.JoinGroup(en4, & net.UDPAddr{IP:group}); err != nil {
        fmt.Println(err)
    }

    // 4.更多的控制
    if err := p.SetControlMessage(ipv4.FlagDst,true); err != nil {
        fmt.Println(err)
    }

    // 5.接收消息
    b := make([]byte,1500)
    for {
        n, cm, src, err := p.ReadFrom(b)
        if err != nil {
            fmt.Println(err)
        }

        if cm.Dst.IsMulticast(){
            if cm.Dst.Equal(group){
                fmt.Printf("received: %s from <%s>\n", b[:n], src)
                n, err = p.WriteTo([]byte("world"), cm, src)
				if err != nil {
					fmt.Println(err)
				}
            }else{
                    fmt.Println("Unknown group")
                    continue
            }
        }
    }

}

// func InterfaceByName(name string) (*Interface, error)
// InterfaceByName返回指定名字的网络接口。

// func ListenPacket(net, laddr string) (PacketConn, error)
// ListenPacket函数监听本地网络地址laddr。网络类型net必须是面向数据包的网络类型：
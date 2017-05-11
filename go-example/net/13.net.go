package main

import (
    "net"
    "os"
    "fmt"
)

func main(){
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr,"Usage: %s hostname\n",os.Args[0])
        fmt.Println("Usage: ",os.Args[0],"hostname")
        os.Exit(1)
    }

    name := os.Args[1]
    addr, err := net.ResolveIPAddr("ip",name)
    if err != nil {
        fmt.Println("Resolution error",err.Error())
        os.Exit(1)
    }
    fmt.Println("Resolved address is ",addr.String())
    os.Exit(0)
}

// go run 13.net.go sg.bi.sensetime.com
// Resolved address is  116.62.37.238

// type IPAddr struct {
//     IP   IP
//     Zone string // IPv6范围寻址域
// }
// IPAddr代表一个IP终端的地址。
// func ResolveIPAddr(net, addr string) (*IPAddr, error)
// ResolveIPAddr将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析。 函数会在参数net指定的网络类型上解析，net必须是"ip"、"ip4"或"ip6"。

// func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
// 这种类型的主要用途是通过IP主机名执行DNS查找。

// ResolveIPAddr函数将对某个主机名执行DNS查询，并返回一个简单的IP地址。
// 然而，通常主机如果有多个网卡，则可以有多个IP地址。它们也可能有多个主机名，作为别名。
// func LookupHost(name string) (cname string, addrs []string, err os.Error)
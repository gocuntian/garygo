package main

import (
    "net"
    "os"
    "fmt"
)

func main(){
    if len(os.Args) != 2{
        fmt.Fprintf(os.Stderr,"Usage: %s ip-addr\n")
        os.Exit(1)
    }

    name := os.Args[1]

    addr := net.ParseIP(name)
    if addr == nil{
        fmt.Println("Invalid address")
    }


    mask := addr.DefaultMask()

    network := addr.Mask(mask)

    ones, bits := mask.Size()
    fmt.Println("Address is: ", addr.String())
    fmt.Println("Mask is (hex): ",mask.String())
    fmt.Println("Network is: ",network.String())
    fmt.Println("Leading ones count is: ", ones)
    fmt.Println("Defaut mask length is:",bits)

    os.Exit(0)
}

// Address is:  127.0.0.1
// Mask is (hex):  ff000000
// Network is:  127.0.0.0
// Leading ones count is:  8
// Defaut mask length is: 32

// func ParseIP(s string) IP
// ParseIP将s解析为IP地址，并返回该地址。如果s不是合法的IP地址文本表示，ParseIP会返回nil。
// 字符串可以是小数点分隔的IPv4格式（如"74.125.19.99"）或IPv6格式（如"2001:4860:0:2001::68"）格式。

// func (ip IP) DefaultMask() IPMask
// 函数返回IP地址ip的默认子网掩码。只有IPv4有默认子网掩码；如果ip不是合法的IPv4地址，会返回nil。

// func (ip IP) Mask(mask IPMask) IP
// Mask方法认为mask为ip的子网掩码，返回ip的网络地址部分的ip。（主机地址部分都置0）

// func (m IPMask) Size() (ones, bits int)
// Size返回m的前导的1字位数和总字位数。如果m不是规范的子网掩码（字位：/^1+0+$/），将返会(0, 0)。
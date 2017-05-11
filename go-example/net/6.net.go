package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "strconv"
)
//需要通过从客户端发送不同的请求来获取不同的时间格式，而且需要一个长连接
func main(){
    service := ":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
    checkError(err)
    listener, err := net.ListenTCP("tcp",tcpAddr)
    checkError(err)
    for {
        conn, err :=listener.Accept()
        if err != nil {
            continue
        }
        go handleClient(conn)
    }
}
// 使用conn.Read()不断读取客户端发来的请求。需要保持与客户端的长连接，所以不能在读取完一次请求后就关闭连接。
// 由于conn.SetReadDeadline()设置了超时，当一定时间内客户端无请求发送，conn便会自动关闭，下面的for循环即会因为连接已关闭而跳出。
// 需要注意的是，request在创建时需要指定一个最大长度以防止flood attack；每次读取到请求处理完毕后，需要清理request，
// 因为conn.Read()会将新读取到的内容append到原内容之后。

// 控制TCP连接
// TCP有很多连接控制函数，平常用到比较多的有如下几个函数：
// func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
// DialTimeout类似Dial但采用了超时。timeout参数如果必要可包含名称解析。

// func (c *IPConn) SetDeadline(t time.Time) error
// SetDeadline设置读写操作绝对期限，实现了Conn接口的SetDeadline方法

// func (c *IPConn) SetReadDeadline(t time.Time) error
// SetReadDeadline设置读操作绝对期限，实现了Conn接口的SetReadDeadline方法

// func (c *IPConn) SetWriteDeadline(t time.Time) error
// SetWriteDeadline设置写操作绝对期限，实现了Conn接口的SetWriteDeadline方法

// func (c *TCPConn) SetKeepAlive(keepalive bool) error
// SetKeepAlive设置操作系统是否应该在该连接中发送keepalive信息
// 设置客户端是否和服务器端保持长连接，可以降低建立TCP连接时的握手开销，对于一些需要频繁交换数据的应用场景比较适用。


func handleClient(conn net.Conn){
    conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
    request := make([]byte,128)
    defer conn.Close()
    for {
        read_len, err :=conn.Read(request)
        if err != nil {
            fmt.Println(err)
            break
        }

        if read_len == 0 {
            break
        }else if string(request) == "timestamp" {
            daytime :=strconv.FormatInt(time.Now().Unix(),10)
            conn.Write([]byte(daytime))
        }else{
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }
        request = make([]byte,128)
    }
}

func checkError(err error){
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
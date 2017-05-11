package main

import (
    "fmt"
    "net"
    "os"
    "time"
)
// 3.net.go 有个缺点，执行的时候是单任务的，不能同时接收多个请求，
// 那么该如何改造以使它支持多并发呢？Go里面有一个goroutine机制

// type Conn interface {
//     // Read从连接中读取数据
//     // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
//     Read(b []byte) (n int, err error)
//     // Write从连接中写入数据
//     // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
//     Write(b []byte) (n int, err error)
//     // Close方法关闭该连接
//     // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
//     Close() error
//     // 返回本地网络地址
//     LocalAddr() Addr
//     // 返回远端网络地址
//     RemoteAddr() Addr
//     // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
//     // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
//     // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
//     // 参数t为零值表示不设置期限
//     SetDeadline(t time.Time) error
//     // 设定该连接的读操作deadline，参数t为零值表示不设置期限
//     SetReadDeadline(t time.Time) error
//     // 设定该连接的写操作deadline，参数t为零值表示不设置期限
//     // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
//     SetWriteDeadline(t time.Time) error
// }
// Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。

// type Listener interface {
//     // Addr返回该接口的网络地址
//     Addr() Addr
//     // Accept等待并返回下一个连接到该接口的连接
//     Accept() (c Conn, err error)
//     // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
//     Close() error
// }

func main(){
    service:=":1200"
    tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
    checkError(err)
    listener,  err := net.ListenTCP("tcp",tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue;
        }
        go handleClient(conn)
    }
}

func handleClient(conn, net.Conn){
    defer conn.Close()
    daytime := time.Now().String()
    conn.Write([]byte(daytime))
}

func checkError(err error){
    if err != nil {
        fmt.Fprintf(os.Stderr,"Fatal error:%s",err.Error())
        os.Exit(1)
    }
}

// 通过把业务处理分离到函数handleClient，实现多并发执行了。
// 增加go关键词就实现了服务端的多并发，可以看出goroutine的强大之处。
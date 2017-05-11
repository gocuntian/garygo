package main

import (
    "fmt"
    "bufio"
    "net"
    "time"
)

var quitSemaphore chan bool

// 一.客户端的代码实现步骤:

// 1)创建一个套接字对象, ip与端口指定到上面我们实现的服务器的ip与端口上.
// 2)使用创建好的套接字对象连接服务器.
// 3)连接成功后, 开启一个goroutine, 在这个goroutine内, 定时的向服务器发送消息, 并接受服务器的返回消息, 直到错误发生或断开连接.

// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
// DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"tcp"、"tcp4"、"tcp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。
func main(){
    var tcpAddr *net.TCPAddr
    tcpAddr, _ = net.ResolveTCPAddr("tcp","127.0.0.1:9999")

    conn, _ := net.DialTCP("tcp",nil,tcpAddr)
    defer conn.Close()
    fmt.Println("connected!")
    go onMessageRecived(conn)

    b := []byte("time\n")
    conn.Write(b)

    <-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn){
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        fmt.Println(msg)
        if err != nil {
            quitSemaphore <- true
            break
        }
        time.Sleep(time.Second)
        b := []byte(msg)
        conn.Write(b)
    }
}


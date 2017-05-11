package main

import (
    "fmt"
    "bufio"
    "net"
)
var quitSemaphore chan bool
// 客户端代码改动
// 客户端代码改动相对简单, 只是加入了用户自己输入聊天信息的功能, 在连接成功并且 启动了消息接收的gorountine后, 加入以下代码:

func main(){
    var tcpAddr *net.TCPAddr
    tcpAddr, _ = net.ResolveTCPAddr("tcp","127.0.0.1:9999")

    conn, _ := net.DialTCP("tcp",nil,tcpAddr)
    defer conn.Close()
    fmt.Println("connected!")
    go onMessageRecived(conn)

    //控制台聊天功能
    for {
         var msg string
         fmt.Scanln(&msg)
         if msg == "quit" {
             break
         }
         b := []byte(msg + "\n")
         conn.Write(b)
    }
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
    }
}
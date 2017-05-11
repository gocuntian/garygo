package main

import (
    "fmt"
    "net"
    "bufio"
)

// 服务端的改动

// 服务器为了实现聊天信息的群体广播, 需要记录所有连接到服务器的客户端信息, 所以, 我们需要添加一个集合来保存所有客户端的连接:

// var ConnMap map[string]*net.TCPConn
// 接着, 每次当有新的客户端连接到服务器时, 需要把这个客户端连接行信息加入集合:

// ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
// 当服务器收到客户端的聊天信息时, 需要广播到所有客户端, 所以我们需要利用上面保存TCPConn的map来遍历所有TCPConn进行广播, 用以下方法实现:

//用来记录所有客户端链接
var ConnMap map[string]*net.TCPConn
func main(){
    var tcpAddr *net.TCPAddr
    ConnMap = make(map[string]*net.TCPConn)
    tcpAddr, _ = net.ResolveTCPAddr("tcp","127.0.0.1:9999")

    tcpListener, _ := net.ListenTCP("tcp",tcpAddr)
    defer tcpListener.Close()

    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }
        fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
        //新链接加入map
        ConnMap[tcpConn.RemoteAddr().String()] = tcpConn
        go tcpPipe(tcpConn)
    }

}

func tcpPipe(conn *net.TCPConn){
    ipStr := conn.RemoteAddr().String()
    defer func(){
        fmt.Println("Disconnected :" + ipStr)
        conn.Close()
    }()
    reader := bufio.NewReader(conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Println(conn.RemoteAddr().String() + ":" + string(message))
        //广播
        boradcastMessage(conn.RemoteAddr().String()+":" + string(message))
    }
}

func boradcastMessage(message string){
    b := []byte(message)
    //遍历所有客户端并发送消息
    for _, conn := range ConnMap {
        conn.Write(b)
    }
}
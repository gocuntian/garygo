package main

import (
    "fmt"
    "time"
    "net"
    "bufio"
)
//  一.服务端的实现思路及步骤:
//  1)创建一个套接字对象, 指定其IP以及端口.
//  2)开始监听套接字指定的端口.
//  3)如有新的客户端连接请求, 则建立一个goroutine, 在goroutine中, 读取客户端消息, 并转发回去, 直到客户端断开连接
//  4)主进程继续监听端口
// type TCPAddr struct {
//     IP   IP
//     Port int
//     Zone string // IPv6范围寻址域
// }

// func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
// ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。
// func (l *TCPListener) Close() error
// Close停止监听TCP地址，已经接收的连接不受影响。
// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
// AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。
// func (c *IPConn) RemoteAddr() Addr
// RemoteAddr返回远端网络地址
func main(){
    var tcpAddr *net.TCPAddr
    
    tcpAddr, _ = net.ResolveTCPAddr("tcp","127.0.0.1:9999")

    tcpListener, _ := net.ListenTCP("tcp",tcpAddr)

    defer tcpListener.Close()

    for {
        tcpConn, err := tcpListener.AcceptTCP()
        if err != nil {
            continue
        }
        fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
        go tcpPipe(tcpConn)
    }
}
// type TCPConn struct {
//     // 内含隐藏或非导出字段
// }
// TCPConn代表一个TCP网络连接，实现了Conn接口。
//func (a *TCPAddr) String() string
// func (c *TCPConn) Close() error
// Close关闭连接
// func (b *Reader) ReadString(delim byte) (line string, err error)
// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
// 如果ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
// 当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
// func (c *TCPConn) Write(b []byte) (int, error)
// Write实现了Conn接口Write方法
func tcpPipe(conn *net.TCPConn){
  ipStr := conn.RemoteAddr().String()
  defer func(){
      fmt.Println("disconnected :" + ipStr)
      conn.Close()
  }()
  //创建一个缓冲读取器
  reader := bufio.NewReader(conn)
  
  for {
      message, err := reader.ReadString('\n')
      if err != nil {
          return
      }
      fmt.Println(string(message))
      msg := time.Now().String() + "\n"
      b := []byte(msg)
      conn.Write(b)
  }
}
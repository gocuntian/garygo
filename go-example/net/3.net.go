package main

import (
    "fmt"
    "net"
    "os"
    "time"
)
//通过net包来创建一个服务器端程序，在服务器端我们需要绑定服务到指定的非激活端口，并监听此端口，当有客户端请求到达的时候可以接收到来自客户端连接的请求。net包中有相应功能的函数
//实现一个简单的时间同步服务，监听7777端口
func main(){
    service := ":7777"
    tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
    checkError(err)
    listener, err := net.ListenTCP("tcp",tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
           continue;
        }

        daytime :=time.Now().String()
        conn.Write([]byte(daytime))
        conn.Close()
    } 
}

// 服务跑起来之后，它将会一直在那里等待，直到有新的客户端请求到达。
// 当有新的客户端请求到达并同意接受Accept该请求的时候他会反馈当前的时间信息。
// 值得注意的是，在代码中for循环里，当有错误发生时，直接continue而不是退出，
// 是因为在服务器端跑代码的时候，当有错误发生的情况下最好是由服务端记录错误，
// 然后当前连接的客户端直接报错而退出，从而不会影响到当前服务端运行的整个服务

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

// type TCPListener struct { /*内含隐藏或非导出字段 */ }
// TCPListener代表一个TCP网络的监听者。使用者应尽量使用Listener接口而不是假设（网络连接为）TCP。

// func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
// ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。

// func (l *TCPListener) Addr() Addr
// Addr返回l监听的的网络地址，一个*TCPAddr。

// func (l *TCPListener) SetDeadline(t time.Time) error
// 设置监听器执行的期限，t为Time零值则会关闭期限限制。

// func (l *TCPListener) Accept() (Conn, error)
// Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。

// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
// AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。

// func (l *TCPListener) Close() error
// Close停止监听TCP地址，已经接收的连接不受影响。

// func (l *TCPListener) File() (f *os.File, err error)
// File方法返回下层的os.File的副本，并将该副本设置为阻塞模式。

// 使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。
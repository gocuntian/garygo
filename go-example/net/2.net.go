package main

import (
    "fmt"
    "io/ioutil"
    "net"
    "os"
)
//模拟一个基于HTTP协议的客户端请求去连接一个Web服务端
func main(){
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr,"Usage : %s host:port",os.Args[0])
        os.Exit(1)
    }

    service :=os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4",service)
    fmt.Println(tcpAddr)//127.0.0.1:80
    checkError(err)
    conn, err := net.DialTCP("tcp",nil,tcpAddr)
    checkError(err)
    fmt.Println(conn)
    fmt.Println(conn.RemoteAddr())

    _,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err :=ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
   // conn.Close()
    os.Exit(0)

}
//通过上面的代码我们可以看出：首先程序将用户的输入作为参数service传入net.ResolveTCPAddr获取一个tcpAddr,然后把tcpAddr传入DialTCP后创建了一个TCP连接conn，通过conn来发送请求信息，最后通过ioutil.ReadAll从conn中读取全部的文本，也就是服务端响应反馈的信息。
// func ReadAll(r io.Reader) ([]byte, error)
// ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。成功的调用返回的err为nil而非EOF。因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。

func checkError(err error){
    if err != nil {
        fmt.Fprintf(os.Stderr,"Fatal err : %s",err.Error())
        os.Exit(1)
    }
}

// type TCPConn struct {
//     // 内含隐藏或非导出字段
// }
// TCPConn代表一个TCP网络连接，实现了Conn接口。

// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
// DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。
// net必须是"tcp"、"tcp4"、"tcp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。

// func (c *TCPConn) LocalAddr() Addr
// LocalAddr返回本地网络地址

// func (c *TCPConn) RemoteAddr() Addr
// RemoteAddr返回远端网络地址

// func (c *TCPConn) SetReadBuffer(bytes int) error
// SetReadBuffer设置该连接的系统接收缓冲

// func (c *TCPConn) SetWriteBuffer(bytes int) error
// SetWriteBuffer设置该连接的系统发送缓冲

// func (c *TCPConn) SetDeadline(t time.Time) error
// SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法

// func (c *TCPConn) SetReadDeadline(t time.Time) error
// SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法

// func (c *TCPConn) SetWriteDeadline(t time.Time) error
// SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法

// func (c *TCPConn) SetKeepAlive(keepalive bool) error
// SetKeepAlive设置操作系统是否应该在该连接中发送keepalive信息

// func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
// SetKeepAlivePeriod设置keepalive的周期，超出会断开

// func (c *TCPConn) SetLinger(sec int) error
// SetLinger设定当连接中仍有数据等待发送或接受时的Close方法的行为。

// 如果sec < 0（默认），Close方法立即返回，操作系统停止后台数据发送；如果 sec == 0，Close立刻返回，操作系统丢弃任何未发送或未接收的数据；如果sec > 0，Close方法阻塞最多sec秒，等待数据发送或者接收，在一些操作系统中，在超时后，任何未发送的数据会被丢弃。

// func (c *TCPConn) SetNoDelay(noDelay bool) error
// SetNoDelay设定操作系统是否应该延迟数据包传递，以便发送更少的数据包（Nagle's算法）。默认为真，即数据应该在Write方法后立刻发送。
//===========================================================/
// func (c *TCPConn) Read(b []byte) (int, error)
// Read实现了Conn接口Read方法

// func (c *TCPConn) Write(b []byte) (int, error)
// Write实现了Conn接口Write方法
//=================================================================//

// func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
// ReadFrom实现了io.ReaderFrom接口的ReadFrom方法

// func (c *TCPConn) Close() error
// Close关闭连接

// func (c *TCPConn) CloseRead() error
// CloseRead关闭TCP连接的读取侧（以后不能读取），应尽量使用Close方法。

// func (c *TCPConn) CloseWrite() error
// CloseWrite关闭TCP连接的写入侧（以后不能写入），应尽量使用Close方法。

// func (c *TCPConn) File() (f *os.File, err error)
// File方法设置下层的os.File为阻塞模式并返回其副本。

// 使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。


// type TCPAddr struct {
//     IP   IP
//     Port int
//     Zone string // IPv6范围寻址域
// }
// TCPAddr代表一个TCP终端地址。

//=========================================================================================
// func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
// ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，
// 解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。
//=========================================================================================
// IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

// func (a *TCPAddr) Network() string
// 返回地址的网络类型，"tcp"。

// func (a *TCPAddr) String() string

//===================================================================================
// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
// DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"tcp"、"tcp4"、"tcp6"；
//如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。
//====================================================================================
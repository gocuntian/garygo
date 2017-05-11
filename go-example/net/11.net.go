package main

import (
    "fmt"
    "net"
    "os"
)
//UDP client
func main(){
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]
    udpAddr, err :=net.ResolveUDPAddr("udp4",service)
    checkError(err)
    conn, err := net.DialUDP("udp",nil,udpAddr)
    checkError(err)
    _, err = conn.Write([]byte("anythingddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd"))
    checkError(err)
    var buf [512]byte
    n, err := conn.Read(buf[0:])
    checkError(err)
    fmt.Println(string(buf[0:n]))
    os.Exit(0)
}

func checkError(err error){
    if err != nil {
        fmt.Fprintf(os.Stderr,"Fatal error ", err.Error())
        os.Exit(1)
    }
}



// Go语言包中处理UDP Socket和TCP Socket不同的地方就是在服务器端处理多个客户端请求数据包的方式不同,
// UDP缺少了对客户端连接请求的Accept函数。其他基本几乎一模一样，只有TCP换成了UDP而已
// type UDPAddr struct {
//     IP   IP
//     Port int
//     Zone string // IPv6范围寻址域
// }
// UDPAddr代表一个UDP终端地址。

// func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
// ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"udp"、"udp4"或"udp6"。

// IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

// func (a *UDPAddr) Network() string
// 返回地址的网络类型，"udp"。

// func (a *UDPAddr) String() string


// type UDPConn struct {
//     // 内含隐藏或非导出字段
// }
// UDPConn代表一个UDP网络连接，实现了Conn和PacketConn接口。

// func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)
// DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"udp"、"udp4"、"udp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。

// func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
// ListenUDP创建一个接收目的地是本地地址laddr的UDP数据包的网络连接。net必须是"udp"、"udp4"、"udp6"；如果laddr端口为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。返回的*UDPConn的ReadFrom和WriteTo方法可以用来发送和接收UDP数据包（每个包都可获得来源地址或设置目标地址）。

// func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
// ListenMulticastUDP接收目的地是ifi接口上的组地址gaddr的UDP数据包。它指定了使用的接口，如果ifi是nil，将使用默认接口。

// func (c *UDPConn) LocalAddr() Addr
// LocalAddr返回本地网络地址

// func (c *UDPConn) RemoteAddr() Addr
// RemoteAddr返回远端网络地址

// func (c *UDPConn) SetReadBuffer(bytes int) error
// SetReadBuffer设置该连接的系统接收缓冲

// func (c *UDPConn) SetWriteBuffer(bytes int) error
// SetWriteBuffer设置该连接的系统发送缓冲

// func (c *UDPConn) SetDeadline(t time.Time) error
// SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法

// func (c *UDPConn) SetReadDeadline(t time.Time) error
// SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法

// func (c *UDPConn) SetWriteDeadline(t time.Time) error
// SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法

// func (c *UDPConn) Read(b []byte) (int, error)
// Read实现Conn接口Read方法

// func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
// ReadFrom实现PacketConn接口ReadFrom方法

// func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
// ReadFromUDP从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。

// ReadFromUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。

// func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
// ReadMsgUDP从c读取一个数据包，将有效负载拷贝进b，相关的带外数据拷贝进oob，返回拷贝进b的字节数，拷贝进oob的字节数，数据包的flag，数据包来源地址和可能的错误。

// func (c *UDPConn) Write(b []byte) (int, error)
// Write实现Conn接口Write方法

// func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
// WriteTo实现PacketConn接口WriteTo方法

// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
// WriteToUDP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。

// WriteToUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。

// func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
// WriteMsgUDP通过c向地址addr发送一个数据包，b和oob分别为包有效负载和对应的带外数据，返回写入的字节数（包数据、带外数据）和可能的错误。

// func (c *UDPConn) Close() error
// Close关闭连接

// func (c *UDPConn) File() (f *os.File, err error)
// File方法设置下层的os.File为阻塞模式并返回其副本。

// 使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。
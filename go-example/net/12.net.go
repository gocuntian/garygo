package main

import (
    "fmt"
    "os"
    "net"
    "time"
)
// type UDPConn struct {
//     // 内含隐藏或非导出字段
// }
// UDPConn代表一个UDP网络连接，实现了Conn和PacketConn接口。
// func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
// ListenUDP创建一个接收目的地是本地地址laddr的UDP数据包的网络连接。net必须是"udp"、"udp4"、"udp6"；如果laddr端口为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。返回的*UDPConn的ReadFrom和WriteTo方法可以用来发送和接收UDP数据包（每个包都可获得来源地址或设置目标地址）。

// func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
// ReadFromUDP从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。

// ReadFromUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。

func main(){
    service := ":1200"
    udpAddr, err := net.ResolveUDPAddr("udp4",service)
    checkError(err)
    conn, err := net.ListenUDP("udp",udpAddr)
    checkError(err)
    for{
        handleClient(conn)
    }
}

func handleClient(conn *net.UDPConn){
    //var buf [512]byte 
    // _, addr, err :=conn.ReadFromUDP(buf[0:]) 
    // Or
    buf := make([]byte,512)
    fmt.Println(buf)
    _, addr, err :=conn.ReadFromUDP(buf)
    if err != nil {
        return
    }
    fmt.Println(buf)
    daytime := time.Now().String()
    conn.WriteToUDP([]byte(daytime),addr)
}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
        os.Exit(1)
    }
}
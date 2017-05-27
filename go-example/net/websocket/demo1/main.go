package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, _ := conn.ReadMessage()
				println(string(msg))
				conn.WriteMessage(mType, msg)
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg))
			}
		}(conn)
	})

	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)
			for range ch {
				conn.WriteJSON(myResponseMsg{
					Username: "xingcuntian",
					Age:      26,
					Sex:      1,
				})
			}
		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					conn.Close()
					return
				}
			}
		}(conn)

		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)
			for range ch {
				conn.WriteJSON(myResponseMsg{
					Username: "xingcuntian",
					Age:      26,
					Sex:      1,
				})
			}
		}(conn)
	})

	http.ListenAndServe(":3000", nil)
}

type myResponseMsg struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}

// var ws = new WebSocket("ws://localhost:3000/v1/ws")
// ws.addEventListener("message",function(e){console.log(e);});
// ws.send("xingcuntian")

// var ws = new WebSocket("ws://localhost:3000/v2/ws")
// ws.send(JSON.stringify({username:"xingcuntian"}))

// var ws = new WebSocket("ws://localhost:3000/v3/ws")
// ws.addEventListener("message",function(e){console.log(e);});
// ws.readyState 1
// ws.OPEN 1
// ws.close()
// ws.readyState 2
// ws.CLOSING	2

// var ws = new WebSocket("ws://localhost:3000/v4/ws")
// ws.close()
// ws.readyState 3
// ws.CLOSING	3

// func ServeFile(w ResponseWriter, r *Request, name string)
// ServeFile回复请求name指定的文件或者目录的内容。
// type Upgrader struct {
// 	//HandshakeTimeout指定握手完成的持续时间。
// 	HandshakeTimeout time.Duration

// 	// ReadBufferSize和WriteBufferSize指定I / O缓冲区大小。
// 	// 如果缓冲区大小为零，则使用HTTP服务器分配的缓冲区。
// 	// I / O缓冲区大小不限制可以发送或接收的消息的大小。
// 	ReadBufferSize, WriteBufferSize int

// 	// 子协议按优先级顺序指定服务器支持的协议。
// 	// 如果此字段已设置，则升级方法通过使用客户端请求的协议选择此列表中的第一个匹配来协商子协议。
// 	Subprotocols []string

// 	// 错误指定生成HTTP错误响应的函数。 如果Error为nil，则使用http.Error来生成HTTP响应。
// 	Error func(w http.ResponseWriter, r *http.Request, status int, reason error)

// 	// 如果请求原始头可以接受，CheckOrigin返回true。
// 	// 如果CheckOrigin为零，则Origin头中的主机不得设置或必须与请求主机匹配。
// 	CheckOrigin func(r *http.Request) bool

// 	//EnableCompression指定服务器是否应尝试协商每个消息压缩（RFC 7692）。
// 	//将此值设置为true不能保证将支持压缩。 目前只支持“无上下文接管”模式。
// 	EnableCompression bool
// }

// func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Conn, error) {
// 升级将HTTP服务器连接升级到WebSocket协议。
// responseHeader包含在响应客户端的升级请求中。 使用responseHeader指定cookie（Set-Cookie）和应用程序协商的子协议（Sec-Websocket-Protocol）。
// 如果升级失败，则升级会使用HTTP错误响应回复客户端。

// func (c *Conn) ReadMessage() (messageType int, p []byte, err error) {
// ReadMessage是一个帮助方法，让读者使用NextReader并从该读取器读取缓冲区。

// func (c *Conn) WriteMessage(messageType int, data []byte) error {
// WriteMessage是一个帮助方法，用于使用NextWriter获取作者，编写消息并关闭作者。

// func (c *Conn) RemoteAddr() net.Addr
// RemoteAddr返回远程网络地址。

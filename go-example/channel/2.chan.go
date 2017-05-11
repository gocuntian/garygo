package main

import (
    "fmt"
    "time"
)
// 无缓冲管道是阻塞
// # 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
// # 这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待
// # 消息 `"ping"`。
func main(){
    message := make(chan string)
    go func(){
        time.Sleep(5 * time.Second)
        message <- "ping"
    }()

    fmt.Println("start") 
    msg := <-message
    fmt.Println(msg)
}
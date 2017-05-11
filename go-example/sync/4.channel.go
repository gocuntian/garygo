package main

import "fmt"
import "time"

// channel阻塞超时处理： 
// goroutine有时候会进入阻塞情况，那么如何避免由于channel阻塞导致整个程序阻塞的发生那？
// 解决方案：通过select设置超时处理

func main(){
    c := make(chan int)
    o := make(chan bool)
    go func(){
        for {
            select {
                case i := <-c:
                  fmt.Println(i)
                ////设置超时时间为３ｓ，如果channel　3s钟没有响应，一直阻塞，则报告超时，进行超时处理．
                case <-time.After(time.Duration(3) * time.Second) :
                fmt.Println("timeout")
                o <- true
                break
            }
        }
    }()
    fmt.Println("out111")
    <- o 
    fmt.Println("out222")
}
package main

import (
    "fmt"
    "runtime"
    "time"
)

// func Stack(buf []byte, all bool) int
// Stack将调用其的go程的调用栈踪迹格式化后写入到buf中并返回写入的字节数。
// 若all为true，函数会在写入当前go程的踪迹信息后，将其它所有go程的调用栈踪迹都格式化写入到buf中。

func main(){
    go func(){
        fmt.Println("i am a goroutine")
        time.Sleep(time.Second)
    }()

    time.Sleep(500 * time.Millisecond)
    buf := make([]byte,1024)
    n :=runtime.Stack(buf,false)
    fmt.Println(string(buf[:n]))
    fmt.Println("=========================")

    n = runtime.Stack(buf,true)
    fmt.Println(string(buf[:n]))
}

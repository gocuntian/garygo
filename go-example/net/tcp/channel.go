package main

import (
    "fmt"
)
// goroutine是一种轻量型的线程, 作为golang语言的语言特性, 可以很简单的在golang中进行多线程的开发. 
// 利用go关键字, 我们能把任何一个方法/函数, 放在一个新的goroutine里执行.

var quit chan bool = make(chan bool)

func main(){
    go testGorountine()
    <-quit
}

func testGorountine() {
    for i := 0; i < 10; i++ {
        fmt.Println("Hello world!")
    }
    quit <- true
}
//终端的输出, 可以看到10行"hello world". 这里, 我们的hello world程序就是利用了gorountine创建了一个多线程/协程程序, 然后利用channel等待开启的协程处理完毕, 才结束主线程.
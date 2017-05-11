package main

import "fmt"
// close(): 
// close主要用来关闭channel通道其用法为close(channel)，并且实在生产者的地方关闭channel，而不是在消费者的地方关闭．
// 并且 关闭channel后，便不可再想channel中继续存入数据，但是可以继续从channel中读取数据．
func main(){
    var ch = make(chan int,20)
    for i := 0; i < 20; i++ {
        ch <- i
    }
    close(ch)
  //  ch <-11 //send on closed channel
    for i := range ch {
        fmt.Println(i)
    }
}
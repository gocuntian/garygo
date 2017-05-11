package main

import "fmt"
//import "time"

func Afunction(ch chan int){
    fmt.Println("finish")
    <-ch
}

func main(){
    ch := make(chan int) //无缓冲的channel
   // ch := make(chan int,3) //有缓冲的channel
    ch <- 1 //  主协程阻塞  死锁
    go Afunction(ch)
   // time.Sleep(2* time.Second) 
}
//首先创建一个无缓冲的channel,　然后在主协程里面向channel　ch 中通过ch<-1命令写入数据，则此时主协程阻塞，
// 就无法执行下面的go Afuntions(ch),自然也就无法解除主协程的阻塞状态，则系统死锁 

// 对于无缓存的channel,放入channel和从channel中向外面取数据这两个操作不能放在同一个协程中，防止死锁的发生；
// 同时应该先利用go 开一个协程对channel进行操作，此时阻塞该go 协程，然后再在主协程中进行channel的相反操作
// （与go 协程对channel进行相反的操作），实现go 协程解锁．即必须go协程在前，解锁协程在后．

// 带缓存channel: 
// 对于带缓存channel，只要channel中缓存不满，则可以一直向 channel中存入数据，直到缓存已满；同理只要channel中缓存不为０，便可以一直从channel中向外取数据，直到channel缓存变为０才会阻塞． 

// 由此可见，相对于不带缓存channel，带缓存channel不易造成死锁，可以同时在一个goroutine中放心使用， 

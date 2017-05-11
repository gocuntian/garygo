package main

//死锁
// func main(){
//     c:=make(chan int)
//     c <- 42
//     val := <-c
//     println(val)
// }

// 死锁. 在这种情况下，两个goroutine互相等待对方释放资源，造成双方都无法继续运行。GO语言可以在运行时检测这种死锁并报错。这个错误是因为锁的自身特性产生的。

// 代码在次以单线程的方式运行，逐行运行。向channel写入的操作（c <- 42）会锁住整个程序的执行进程，
// 因为在同步channel中的写操作只有在读取器准备就绪后才能成功执行。然而在这里，我们在写操作的下一行才创建了读取器。
// 为了使程序顺利执行，需要做如下改动：
func main(){
    c:=make(chan int)
    go func(){
        c<-42
    }()
    val := <-c
    println(val)
}
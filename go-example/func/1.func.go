package main

import (
    "fmt"
)
// 总结下Go语言中的错误处理吧。对于初学者来说很重要。
// Go语言中延迟函数defer充当着 try…catch 的重任，使用起来也非常简便，然而在实际应用中，
// 很多gopher并没有真正搞明白defer、return和返回值之间的执行顺序，从而掉进坑中 。

// 匿名返回值的情况：
func main(){
    fmt.Println("a return :", a()) //a return : 0
}

func a() int {
    var i int
    defer func(){
        i++
        fmt.Println("a defer2:",i) //a defer2: 2
    }()

    defer func(){
        i++
        fmt.Println("a defer1:",i) //a defer1: 1
    }()
    return i
}
// a defer1: 1
// a defer2: 2
// a return : 0
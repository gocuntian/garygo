package main

import (
    "fmt"
)
//非匿名返回值的情况:
func main(){
    fmt.Println("b return :",b()) // b return : 2
}

func b() (i int) {
    defer func(){
        i++
        fmt.Println("b defer2:",i) //b defer2: 2
    }()
    defer func(){
        i++
        fmt.Println("b defer1:",i) // b defer1: 1
    }()
    return i //或者直接return 效果一样
}
// b defer1: 1
// b defer2: 2
// b return : 2
package main

import (
    "fmt"
)

func main(){
    c:=c()
    fmt.Println("c return :",*c, c) // c return : 2 0xc42000e2d8
}

func c() *int {
    var i int
    defer func(){
        i++
        fmt.Println("c defer2:",i,&i) // c defer2: 2 0xc42000e2d8
    }()
    defer func(){
        i++
        fmt.Println("c defer1:",i,&i) // c defer1: 1 0xc42000e2d8
    }()
    return &i
}
// c defer1: 1 0xc42000e2d8
// c defer2: 2 0xc42000e2d8
// c return : 2 0xc42000e2d8
// 虽然 c()*int 的返回值没有被提前声明，但是由于 c()*int 的返回值是指针变量，那么在return将变量 i 的地址赋给返回值后，
// defer再次修改了 i 在内存中的实际值，因此return调用RET退出函数时返回值虽然依旧是原来的指针地址，但是其指向的内存实际值已经被成功修改了。
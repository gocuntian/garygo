package main

import "fmt"

// 数组作为参数时，函数内部不改变数组内部的值，除非是传入数组的指针。
// 数组的指针：*[3]int
// 指针数组：[2]*int

func main(){
    //可以用new创建数组，并返回 *数组的指针*
    var a = new([5]int)
    test(a)
    fmt.Println(a,len(a)) 
}

func test(a *[5]int){
    a[1] = 5
}

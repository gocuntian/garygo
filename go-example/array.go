package main

import "fmt"
// 数组是一个具有相同数据类型的元素组成的固定长度的有序集合。
// 在Go语言中，数组是值类型，长度是类型的组成部分，也就是说"[10]int"和“[20]int”是完全不同的两种数组类型。
// 同类型的两个数组支持"=="和"!="比较，但是不能比较大小。
// 数组作为参数时，函数内部不改变数组内部的值，除非是传入数组的指针。
// 数组的指针：*[3]int
// 指针数组：[2]*int
func main(){
    //创建了一个具有5个元素的整型数组
    // 元素的数据类型和数组长度都是数组的一部分
    // 默认情况下，数组元素都是零值 对于整数，零值就是0
    var a [5]int
    fmt.Println("emp: ",a)
    //使用索引来设置数组元素的值，就像这样 "array[index] = value"  或者使用索引来获取元素值，
    // 就像这样"array[index]"
    a[4] = 100
    fmt.Println("set: ",a)
    fmt.Println("get: ",a[4])

    // 内置的len函数返回数组的长度
    fmt.Println("len: ",len(a))

    //同时定义和初始化一个数组
    b:=[5]int{0,0,2,0,100}
    fmt.Println("dcl: ", b)

    if a == b {
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }


    // 数组都是一维的，但是你可以把数组的元素定义为一个数组 来获取多维数组结构
    var twoD [2][3]int
    for i:=0;i<2;i++{
        for j:=0;j<3;j++{
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    //拥有固定长度是数组的一个特点，但是这个特点有时候会带来很多不便，尤其在一个集合元素个数不固定的情况下。这个时候我们更多地使用切片。

}
package main

import (
    "fmt"
    "flag"
)

var bFlog = flag.Bool("b",false,"If show message")

func main(){
    //第一个参数，为参数名称，第二个参数为默认值，第三个参数是说明
    name :=flag.String("name","xingcuntian","input your name.")
    age :=flag.Int("age",0,"input your age.")
    flag.Parse()
    // 来解析命令行参数写入注册的flag里。
    // 解析之后，flag的值可以直接使用。如果你使用的是flag自身，它们是指针；如果你绑定到了某个变量，它们是值。
    if !*bFlog{
        fmt.Println("Mr. Watson, Come Here, I Want You!")
    }else{
        fmt.Println("Hello, ", *name, " age:",*age)
    }
}
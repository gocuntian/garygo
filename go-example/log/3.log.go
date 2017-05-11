package main

import (
    "fmt"
    "log"
)
//对于log.Panic接口，该函数把日志内容刷到标准错误后调用 panic 函数
// func Panicf(format string, v ...interface{})
// Panicf等价于{Printf(v...); panic(...)}

// func Panic(v ...interface{})
// Panic等价于{Print(v...); panic(...)}

// func Panicln(v ...interface{})
// Panicln等价于{Println(v...); panic(...)}

func test_deferpanic(){
    defer func(){
        fmt.Println("--first--")
        if err :=recover(); err != nil {
            fmt.Println(err)
        }
    }()
    log.Panicln("test for defer Panic")
    defer func(){
        fmt.Println("--second--")
    }()
}

func main(){
    test_deferpanic()
}

// 2017/03/22 18:59:47 test for defer Panic
// --first--
// test for defer Panic

// 首先输出了“test for defer Panic”，然后第一个defer函数被调用了并输出了“--first--”，
// 但是第二个defer 函数并没有输出，可见在Panic之后声明的defer是不会执行的
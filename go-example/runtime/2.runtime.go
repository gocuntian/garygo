package main

import (
    "fmt"
    "runtime"
)

// func GOROOT() string
// GOROOT返回Go的根目录。如果存在GOROOT环境变量，返回该变量的值；否则，返回创建Go时的根目录。

// func Version() string
// 返回Go的版本字符串。它要么是递交的hash和创建时的日期；要么是发行标签如"go1.3"。

// func NumCPU() int
// NumCPU返回本地机器的逻辑CPU个数。

// func GOMAXPROCS(n int) int
// GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉。
func main(){
    fmt.Println(runtime.GOROOT())///usr/local/go
    fmt.Println(runtime.Version())//go1.8
    fmt.Println(runtime.NumGoroutine())//1
    fmt.Println(runtime.NumCPU())//8
    fmt.Println(runtime.GOMAXPROCS(-1))//8
}





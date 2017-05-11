package main

import (
    "fmt"
    "time"
)

//无缓冲channel在读和写是都会阻塞
func Afunction(ch chan int){
    fmt.Println("finish")
    time.Sleep(2 * time.Second)
    <-ch
}

func main(){
    ch := make(chan int) //初始化无缓冲的channel
    go Afunction(ch)
    ch <- 1
}

//首先创建一个无缓冲channel　ch,　然后执行　go Afuntion(ch),此时执行＜-ch，则Afuntion这个函数便会阻塞，
//不再继续往下执行，直到主进程中ch<-1向channel　ch 中注入数据才解除Afuntion该协程的阻塞．
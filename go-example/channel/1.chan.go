package main

import (
    "fmt"
    "time"
)

//定义一个结构
type Ball struct{
    hits int
}

func main(){
   


    // //创建一个可传输Ball的channel
    // table := make(chan *Ball)
    // //分别启动ping/pong的goroutine
    //  go Player("Ping",table)
    //  go Player("Pong",table)
    // //  // 一个Ball进入channel,开始游戏
    //  table <- new(Ball)
    // //  //“主”程序暂停1s,等待ping/pong的goroutine执行
	//  time.Sleep(1 * time.Second)
    //  //从channel取出Ball，游戏开始
	// <-table



}

func Player(name string, table chan *Ball){
    for {
        //channel取出Ball,并hits++
        ball:= <-table
        fmt.Println(*ball)
        ball.hits++
        fmt.Println(ball.hits)
        fmt.Println(name,ball.hits)
        //暂停1ms
        time.Sleep(1 * time.Millisecond)
        //将Ball放回channel
        table <- ball
    }
}
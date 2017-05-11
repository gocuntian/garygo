package main

import (
    "fmt"
    "runtime"
)
//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
//Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。


func main(){
    go say("world")
    say("hello")
}

func say(s string){
    for i :=0;i < 5; i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

// hello
// world
// hello
// world
// world
// hello
// world
// hello
// world
// hello
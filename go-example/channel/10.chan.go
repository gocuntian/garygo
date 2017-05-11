package main

import (
    "fmt"
    "time"
    "strconv"
)
// Channels and select
// golang 的 select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。
// 注意到 select 的代码形式和 switch 非常相似， 不过 select 的 case 里的操作语句只能是  IO 操作

func makeCakeAndSend(cs chan string, flavor string, count int) {
    for i :=1; i <= count; i++ {
        cakeName :=flavor + "Cake " + strconv.Itoa(i)
        cs <- cakeName
    }
    close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string) {
    strbry_closed, choco_closed := false, false
    for {
        if (strbry_closed && choco_closed) { 
             fmt.Println("channel closed!")
                return 
            }
        fmt.Println("Waring for new cake ...")
        select {
            case cakeName, strbry_ok := <-strbry_cs:
                if (!strbry_ok) {
                    strbry_closed = true
                   // fmt.Println("... Strawberry channel closed!")
                }else{
                    fmt.Println("Received from Strawberry channel. Now packing",cakeName)
                }
            case cakeName, choco_ok := <-choco_cs:
                if (!choco_ok) {
                    choco_closed = true
                   // fmt.Println("... Chocolate channel closed!")
                }else{
                    fmt.Println("Received from Chocolate channel. Now packing",cakeName)
                }    
        }
    }
}

func main(){
    strbry_cs := make(chan string)
    choco_cs := make(chan string)

    go makeCakeAndSend(choco_cs,"Chocolate",3)
    go makeCakeAndSend(strbry_cs, "Strawberry",3)

    go receiveCakeAndPack(strbry_cs,choco_cs)

    time.Sleep(2 * 1e9)
}
//select 会一直等待等到某个 case 语句完成， 也就是等到成功从 ch1 或者 ch2 中读到数据。 则 select 语句结束
package main

import (
    "fmt"
    "time"
    "strconv"
)
//Channels and range
// Go提供了range关键词,当它与Channel 一起使用的时候他会等待channel的关闭。
func makeCakeAndSend(cs chan string, count int){
    for i := 1; i <= count; i++ {
        cakeName := "Strawberry Cake "+ strconv.Itoa(i)
        cs <- cakeName
    }
}

func receiveCakeAndPack(cs chan string){
    for s :=range cs{
        fmt.Println("Packing received cake: ",s)
    }
}

func main(){
    cs := make(chan string)
    go makeCakeAndSend(cs,6)
    go receiveCakeAndPack(cs)

    time.Sleep(3 * 1e9)
}
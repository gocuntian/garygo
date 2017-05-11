package main

import (
    "fmt"
    "time"
)

func main(){
    timer := time.NewTimer(time.Second * 2)

    <-timer.C

    fmt.Println("Timer 1 expired")

    timer1 := time.NewTimer(time.Second)

    go func(){
        <-timer1.C
        fmt.Println("Timer 2 expired")
    }()

    stop1 := timer1.Stop()
    if stop1 {
        fmt.Println("Timer 2 stopped")
    }
}
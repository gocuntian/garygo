package main

import (
    "fmt"
    "time"
)

func main(){
    message:=make(chan string)
    count := 3
    go func(){
        for i:=1; i<=count; i++ {
            fmt.Println("send messge")
            message <- fmt.Sprintf("messge %d",i)
        }
    }()

    time.Sleep(time.Second * 3)
    for i:=1; i<=count; i++ {
        fmt.Println(<-message)
    }
}
// send messge
// messge 1
// send messge
// send messge
// messge 2
// messge 3
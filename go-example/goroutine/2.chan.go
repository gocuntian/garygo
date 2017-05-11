package main

import (
    "fmt"
    "time"
)

func main(){
    c := make(chan bool)
    go func(){
        time.Sleep(1 * time.Second)
        fmt.Println("Please go!")
        <-c
        time.Sleep(1 * time.Second)
    }()
    c <-true
}
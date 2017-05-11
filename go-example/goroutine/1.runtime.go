package main

import (
    "fmt"
    "runtime"
)

func main(){
    runtime.GOMAXPROCS(runtime.NumCPU())
    fmt.Println(runtime.NumCPU())
    c := make(chan bool, 10)
    for i:= 0; i < 10; i++ {
        go Go(c,i)
    }

    for i := 0; i < 10; i++ {
        <-c
    }
}

func Go(c chan bool, index int) {
    a := 1
    for i:=0; i < 1000000; i++ {
        a += 1
    }
    fmt.Println(index, a)
    c <- true
}
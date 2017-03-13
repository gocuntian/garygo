package main

import (
    "fmt"
    "time"
)

func DoTickerWork(res chan interface{} , timeout <-chan time.Time){
    t:=time.NewTicker(3 * time.Second)
    go func(){
        defer close(res)
        i:=1
        for{
            <-t.C 
            fmt.Printf("start %d th worker\n",i)
            res <-i
            i++
        }
    }()
    <-timeout
    t.Stop()
    return
}

func main(){
    res:=make(chan interface{},10000)
    timeout:=time.After(10 * time.Second)
    DoTickerWork(res,timeout)
    for v:=range res{
        fmt.Println(v)
    }
}
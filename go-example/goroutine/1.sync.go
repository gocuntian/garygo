package main

import (
    "log"
    "sync"
    "time"
)

func main(){
    var wg sync.WaitGroup
    wg.Add(1)
    go waitTime(&wg)
    wg.Wait()
}

func waitTime(wg *sync.WaitGroup) {
    time.Sleep(time.Second * 2)
    log.Println("wait finished!")
    wg.Done()
}
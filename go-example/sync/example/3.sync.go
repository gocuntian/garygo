package main

import (
    "fmt"
    "sync"
    "time"
)

type mutexCounter struct{
    mx sync.Mutex
    x int64
}

func (c *mutexCounter) Add(x int64) {
    c.mx.Lock()
    c.x += x
    c.mx.Unlock()
}

func(c *mutexCounter) Value() (x int64) {
    c.mx.Lock()
    x = c.x
    c.mx.Unlock()
    return
}

func main(){
    counter := mutexCounter{}
    for i := 0; i < 100; i++ {
        go func(no int) {
            for i := 0; i < 10000; i++ {
                counter.Add(1)
            }
        }(i)
    }
    time.Sleep(time.Second)
    fmt.Println(counter.Value())
}
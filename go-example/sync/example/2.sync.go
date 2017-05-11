package main

import (
    "fmt"
    "runtime"
    "sync/atomic"
    "time"
)

type atomicCounter struct{
    val int64
}

func (c *atomicCounter) Add(x int64) {
    atomic.AddInt64(&c.val,x)
    runtime.Gosched()
}

//func Gosched()
//Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。

func(c *atomicCounter) Value() int64 {
    return atomic.LoadInt64(&c.val)
}

func main(){
    counter := atomicCounter{}
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
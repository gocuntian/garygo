package main

import (
    "fmt"
    "sync"
    "time"
)
//在限速时，一种方案是丢弃请求，即请求速度太快时，对后进入的请求直接抛弃。

type LimitRate struct{
    rate  int
    begin time.Time
    count int
    lock  sync.Mutex
}

func (l *LimitRate) Limit() bool{
    result := true
    l.lock.Lock()
    if l.count == l.rate {
        if time.Now().Sub(l.begin) >= time.Second {
            l.begin = time.Now()
            l.count = 0
        }else{
            result = false
        }
    }else{
        l.count++
    }
    l.lock.Unlock()
    return result
}
//SetRate 设置每秒允许的请求数
func (l *LimitRate) SetRate(r int){
    l.rate = r
    l.begin = time.Now()
}
//GetRate 获取每秒允许的请求数
func (l *LimitRate) GetRate() int {
    return l.rate
}

func main(){
    var wg sync.WaitGroup
    var lr LimitRate
    lr.SetRate(10)
    for i:=0; i<100;i++ {
        wg.Add(1)
            go func(){
                if lr.Limit() {
                    fmt.Println("Got it!",i)
                }
                wg.Done()
            }()
    }
    wg.Wait()
}
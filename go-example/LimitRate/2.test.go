package main

import (
    "fmt"
    "time"
    "sync"
)
//在限速时，另一种方案是等待，即请求速度太快时，后到达的请求等待前面的请求完成后才能运行。这种方案类似一个队列。
type LimitRate struct{
    rate       int
    interval   time.Duration
    lastAction time.Time
    lock       sync.Mutex
}

func (l *LimitRate) Limit() bool {
    result :=false
    for {
        l.lock.Lock()
        //判断最后一次执行的时间与当前的时间间隔是否大于限速速率
        if time.Now().Sub(l.lastAction) > l.interval {
            l.lastAction = time.Now()
            result = true
        }
        l.lock.Unlock()
        if result {
            return result
        }
        time.Sleep(l.interval)
    }
}

func (l *LimitRate) SetRate(r int) {
    l.rate = r
    l.interval = time.Microsecond * time.Duration(1000 * 1000/l.rate)
}

func (l *LimitRate) GetRate() int {
    return l.rate
}

func main() {
    var wg sync.WaitGroup
    var lr LimitRate
    lr.SetRate(10)
    
    b:=time.Now()
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            if lr.Limit() {
                fmt.Println("Got it!")
            }
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(time.Since(b))
}
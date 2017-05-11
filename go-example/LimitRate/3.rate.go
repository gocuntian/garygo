package main

import (
    "fmt"
    "time"
    "golang.org/x/time/rate"
)

func main() {
    l := rate.NewLimiter(10, 10)
    for {
        r := l.ReserveN(time.Now(), 1)
        time.Sleep(r.Delay())
        fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
    }
}
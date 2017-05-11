package main

import (
    "fmt"
    "time"
    "golang.org/x/net/context"
    "golang.org/x/time/rate"
)

func main(){
    l := rate.NewLimiter(1,1)
    c, _ := context.WithCancel(context.TODO())
    fmt.Println(l.Limit(),l.Burst())
    for {
        l.Wait(c)
       // time.Sleep(200 * time.Millisecond)
        fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
    }
}
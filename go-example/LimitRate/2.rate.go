package main

import (
    "fmt"
    "time"
    "golang.org/x/time/rate"
)

func main(){
    l := rate.NewLimiter(1000,1000)
     for {
        if l.AllowN(time.Now(), 1) {
            fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
        } else {
            time.Sleep(6     * time.Second)
        }
    }
}

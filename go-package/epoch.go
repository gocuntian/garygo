package main

import (
    "fmt"
    "time"
)

func main(){
    now:=time.Now()
    secs:=now.Unix()
    nanos:=now.UnixNano()
    fmt.Println(now)
    //2017-01-15 00:12:55.662339871 +0800 CST
    millis:=nanos / 1000000

    fmt.Println(secs)
    //1484410375
    fmt.Println(millis)
    //1484410375662
    fmt.Println(nanos)
    //1484410375662339871
    fmt.Println(time.Unix(secs,0))
    //2017-01-15 00:12:55 +0800 CST
    fmt.Println(time.Unix(0,nanos))
    //2017-01-15 00:12:55.662339871 +0800 CST
}
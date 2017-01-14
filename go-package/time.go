package main

import (
    "fmt"
    "time"
)

func main(){
    p:=fmt.Println

   

    then:=time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    
    p(then)

    p(then.Year())
    p(then.Month())
fmt.Println("=====================================>\r\n")
    now:=time.Now();
    p(now)    
    p(now.Year())
    p(now.Month())
    p(now.Day())
    p(now.Hour())
    p(now.Minute())
    p(now.Second())
    p(now.Nanosecond())
    p(now.Location())
    p(now.Weekday())
fmt.Println("=====================================>\r\n")

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

fmt.Println("=====================================>\r\n")
    diff:=now.Sub(then)
    p(diff)
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))




}

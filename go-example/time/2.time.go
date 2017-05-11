package main

import (
    "fmt"
    "time"
)
// type Time struct {
//     // 内含隐藏或非导出字段
// }
//func (t Time) Date() (year int, month Month, day int)
//func Now() Time

func main(){
    year, month, day :=time.Now().Date()
    if month == time.March{
        fmt.Println("this is ",month)
    }
    fmt.Println(year,month,day)
    fmt.Println(time.Now().Month().String())
    fmt.Println(time.Now().Weekday().String())

    t := time.Date(2017, time.March,20,15,0,0,0,time.UTC)
    fmt.Printf("Go launched at %s\n", t.Local())
}
package main

import (
    "fmt"
    "time"
)

// 当前时间 time.Now()
// 把时间格式化成字符串(time->string) : time.Now().Format("2006-01-02 15:04:05")
// 把日期字符串转化为时间(string -> time) : time.Parse("01-02-2006", "06-17-2013")
// 把纳秒转化为时间字符串(int64 -> string): time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
// 获取当前时间纳秒时间戳 time.Now().UnixNano()
// 自己组装时间 time.Date() (看下面)

func main(){
    StringToTime()
    TimeFormat()
}

func StringToTime(){
    layout := "2006-01-02 15:04:05"
    str := "2016-07-25 11:45:26"
    t, err := time.Parse(layout, str)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(t)
}

func TimeFormat(){
    const layout = "Jan 2, 2006 at 3:04pm (MST)"
    t := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
    fmt.Println(t.Format(layout))
    fmt.Println(t.UTC().Format(layout))
}


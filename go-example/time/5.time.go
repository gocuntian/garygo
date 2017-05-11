package main

import (
    "fmt"
    "time"
)
// func (t Time) Unix() int64
// Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）

// func Unix(sec int64, nsec int64) Time
// Unix创建一个本地时间，对应sec和nsec表示的Unix时间（从January 1, 1970 UTC至该时间的秒数和纳秒数）。
// nsec的值在[0, 999999999]范围外是合法的。

// func (t Time) Format(layout string) string
// Format根据layout指定的格式返回t代表的时间点的格式化文本表示

// func Parse(layout, value string) (Time, error)
// Parse解析一个格式化的时间字符串并返回它代表的时间


func main(){
    //获取时间戳
    timestamp := time.Now().Unix()
    fmt.Println(timestamp)

    //格式化为字符串，tm 为 Time类型
    tm := time.Unix(timestamp,0)
    fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
    fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))

    //从字符串转为时间戳，第一个参数是格式，第二个参数是要转换的时间字符串
    tm2, _ := time.Parse("01/02/2006",tm.Format("01/02/2006"))
    fmt.Println(tm2.Unix())
}
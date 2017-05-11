package main

import (
    "fmt"
    "time"
)

// 把时间原点设计在了 2006-01-02 15:04:05

// 其实人家这个日期是有意义的：

// 2006-01-02T15:04:05Z07:00
// 每个数字的意义：

// 1 2 3 4 5 6 7
        
// 月 日 时 分 秒 年 时 区

func main(){
    // 1 toTime格式化（Parsing）：
    // 用法一： 使用毫秒数

    t:= time.Unix(1362984425, 0)
    nt := t.Format("2006-01-02 15:04:05")
    fmt.Println(nt)

     //用法二： 使用固定的字符串进行格式化获得time 对象
     const TimeFormat = "2006-01-02 15:04:05"
     t2, err:=time.Parse(TimeFormat,"2013-03-11 14:47:05")
     if err!=nil{
         fmt.Print(err)
     }
     fmt.Println(t2)



}
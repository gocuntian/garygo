package main

import (
    "log"
)
// const (
//     // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
//     // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
//     Ldate         = 1 << iota     // 日期：2009/01/23
//     Ltime                         // 时间：01:23:23
//     Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
//     Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
//     Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
//     LstdFlags     = Ldate | Ltime // 标准logger的初始值
// )
func main(){
    Ldefault()
    Ldate()
    Ltime()
    Lmicroseconds()
    Llongfile()
    Lshortfile()
    LUTC()
}
// func (l *Logger) SetFlags(flag int)
// SetFlags设置logger的输出选项。
func Ldefault(){
    log.Println("这是默认的格式\n")
}

func Ldate(){
    log.SetFlags(log.Ldate)
    log.Println("这是输出日期格式\n")
}

func Ltime(){
    log.SetFlags(log.Ltime)
    log.Println("这是输出时间格式\n")
}

func Lmicroseconds(){
    log.SetFlags(log.Lmicroseconds)
    log.Println("这是输出微秒格式\n")
}

func Llongfile(){
    log.SetFlags(log.Llongfile)
    log.Println("这是输出路径+文件名+行号格式\n")
}

func Lshortfile(){
    log.SetFlags(log.Lshortfile)
    log.Println("这是输出文件名+行号格式\n")
}

func LUTC(){
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
    log.Println("这是输出 使用标准的UTC时间格式 格式\n")
}
// 2017/03/23 17:20:22 这是默认的格式

// 2017/03/23 这是输出日期格式

// 17:20:22 这是输出时间格式

// 17:20:22.797905 这是输出微秒格式

// /data/golang/src/github.com/xingcuntian/go_test/go-example/log/5.log.go:48: 这是输出路径+文件名+行号格式

// 5.log.go:53: 这是输出文件名+行号格式

// 2017/03/23 09:20:22.797925 这是输出 使用标准的UTC时间格式 格式




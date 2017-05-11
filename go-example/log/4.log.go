package main

import (
    "os"
    "log"
)

// type Logger struct { }
// Logger类型表示一个活动状态的记录日志的对象，它会生成一行行的输出写入一个io.Writer接口。
// 每一条日志操作会调用一次io.Writer接口的Write方法。
// Logger类型的对象可以被多个线程安全的同时使用，它会保证对io.Writer接口的顺序访问。

// func New(out io.Writer, prefix string, flag int) *Logger
// New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。


// （1）输出位置out，是一个io.Writer对象，该对象可以是一个文件也可以是实现了该接口的对象。通常我们可以用这个来指定日志输出到哪个文件。
// （2）prefix 我们在前面已经看到，就是在日志内容前面的东西。我们可以将其置为 "[Info]" 、 "[Warning]"等来帮助区分日志级别。
// （3） flags 是一个选项，显示日志开头的东西，可选的值有：

// Ldate         = 1 << iota     // 形如 2009/01/23 的日期
// Ltime                         // 形如 01:23:23   的时间
// Lmicroseconds                 // 形如 01:23:23.123123   的时间
// Llongfile                     // 全路径文件名和行号: /a/b/c/d.go:23 
// Lshortfile                    // 文件名和行号: d.go:23
// LstdFlags     = Ldate | Ltime // 日期和时间

func main(){
    fileName := "Info_First.log"
    logFile, err := os.Create(fileName)
    defer logFile.Close()
    if err != nil{
        log.Fatalln("create file error")
    }
    debugLog := log.New(logFile,"[Info]",log.Llongfile)
    debugLog.Println("A Info message here")
    debugLog.SetPrefix("[Debug]")
    debugLog.Println("A Info message here ")
}
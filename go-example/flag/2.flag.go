package main

import (
    "fmt"
    "flag"
)

func main(){
    //func StringVar(p *string, name string, value string, usage string)
   // StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量。
    var svar string

	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()
    fmt.Println("svar:", svar)
}
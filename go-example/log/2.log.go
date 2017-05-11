package main

import (
    "fmt"
    "log"
)
//  log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。
//  但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用，
func test_deferfatal(){
    defer func(){
        fmt.Println("--first--")
    }()
    log.Fatalln("test for defer Fatal")
}

func main(){
    test_deferfatal()
}
// 2017/03/22 18:51:26 test for defer Fatal
// exit status 1

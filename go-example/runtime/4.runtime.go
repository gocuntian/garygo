package main

import (
    "fmt"
    "runtime"
)
//runtime.Caller 的用法
// func Caller(skip int) (pc uintptr, file string, line int, ok bool)
// runtime.Caller 返回当前 goroutine 的栈上的函数调用信息. 主要有当前的 pc 值和调用的文件和行号等信息. 若无法获得信息, 返回的 ok 值为 false.

// 其输入参数 skip 为要跳过的栈帧数, 若为 0 则表示 runtime.Caller 的调用者.

// 注意:由于历史原因, runtime.Caller 和 runtime.Callers 中的 skip 含义并不相同, 后面会讲到.

func main(){
    fun1()
}

func fun1(){
    for skip :=0; ; skip++ {
        pc, file, line, ok := runtime.Caller(skip)
        if !ok{
            break
        }
        fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
    }
}
// skip = 0, pc = 4706783, file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/4.runtime.go, line = 21
// skip = 1, pc = 4706704, file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/4.runtime.go, line = 16
// skip = 2, pc = 4356714, file = /usr/local/go/src/runtime/proc.go, line = 185
// skip = 3, pc = 4510609, file = /usr/local/go/src/runtime/asm_amd64.s, line = 2197

// 其中 skip = 0 为当前文件的 main.main 函数, 以及对应的行号.

// 另外的 skip = 1 和 skip = 2 也分别对应2个函数调用. 通过查阅 runtime/proc.c 文件的代码, 我们可以知道对应的函数分别为 runtime.main 和 runtime.goexit.

// 整理之后可以知道, Go的普通程序的启动顺序如下:

// runtime.goexit 为真正的函数入口(并不是main.main)
// 然后 runtime.goexit 调用 runtime.main 函数
// 最终 runtime.main 调用用户编写的 main.main 函数
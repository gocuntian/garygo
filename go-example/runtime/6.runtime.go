package main

import (
    "fmt"
    "runtime"
)
//runtime.Callers 和 runtime.Caller 的异同

func main(){
    fun1()
}

func fun1() {
    for skip := 0; ; skip++ {
        pc, file, line, ok := runtime.Caller(skip)
        if !ok {
            break
        }
        fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
    }

    pc := make([]uintptr, 1024)
    for skip := 0; ; skip++ {
        n := runtime.Callers(skip, pc)
        if n <= 0 {
            break
        }
        fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
    }
}
// skip = 0, pc = 4706914, file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/6.runtime.go, line = 15
// skip = 1, pc = 4706832, file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/6.runtime.go, line = 10
// skip = 2, pc = 4356842, file = /usr/local/go/src/runtime/proc.go, line = 185
// skip = 3, pc = 4510737, file = /usr/local/go/src/runtime/asm_amd64.s, line = 2197
// skip = 0, pc = [4220961 4707391 4706832 4356842 4510737]
// skip = 1, pc = [4707391 4706832 4356842 4510737]
// skip = 2, pc = [4706832 4356842 4510737]
// skip = 3, pc = [4356842 4510737]
// skip = 4, pc = [4510737]

// 比如输出结果可以发现, 4706832 4356842 4510737 这个 pc 值是相同的. 
// 它们分别对应 main.main, runtime.main 和 runtime.goexit 函数.

// runtime.Caller 输出的 4706914 和 runtime.Callers 输出的 4707391 并不相同. 
// 这是因为, 这两个函数的调用位置并不相同, 因此导致了 pc 值也不完全相同.

// 最后就是 runtime.Callers 多输出一个 4220961 值, 对应runtime.Callers内部的调用位置.
package main

import (
    "fmt"
    "runtime"
)

// runtime.FuncForPC 的用途
// func runtime.FuncForPC(pc uintptr) *runtime.Func
// func (f *runtime.Func) FileLine(pc uintptr) (file string, line int)
// func (f *runtime.Func) Entry() uintptr
// func (f *runtime.Func) Name() string
// 其中 runtime.FuncForPC 返回包含给定 pc 地址的函数, 如果是无效 pc 则返回 nil .
// runtime.Func.FileLine 返回与 pc 对应的源码文件名和行号. 安装文档的说明, 如果pc不在函数帧范围内, 则结果是不确定的.
// runtime.Func.Entry 对应函数的地址. 
// runtime.Func.Name 返回该函数的名称.

func main(){
    fun1()
}

func fun1(){
    for skip :=0; ; skip++{
        pc, _, _, ok := runtime.Caller(skip)
        if !ok {
            break
        }
        p :=runtime.FuncForPC(pc)
        file, line :=p.FileLine(0)
        fmt.Printf("skip = %v, pc = %v\n", skip, pc)
        fmt.Printf("  file = %v, line = %d\n", file, line)
        fmt.Printf("  entry = %v\n", p.Entry())
        fmt.Printf("  name = %v\n", p.Name())
    }

// skip = 0, pc = 4706898
//   file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/7.runtime.go, line = 21
//   entry = 4706832
//   name = main.fun1
// skip = 1, pc = 4706816
//   file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/7.runtime.go, line = 17
//   entry = 4706784
//   name = main.main
// skip = 2, pc = 4356714
//   file = /usr/local/go/src/runtime/proc.go, line = 106
//   entry = 4356192
//   name = runtime.main
// skip = 3, pc = 4510721
//   file = /usr/local/go/src/runtime/asm_amd64.s, line = 2197
//   entry = 4510720
//   name = runtime.goexit

    fmt.Println("-------------------------")

 pc :=make([]uintptr,1024)
 for skip :=0; ; skip++ {
     n :=runtime.Callers(skip,pc)
     if n <= 0 {
         break
     }
     fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])

     for j :=0; j < n; j++ {
         p := runtime.FuncForPC(pc[j])
         file, line := p.FileLine(0)

        fmt.Printf("  skip = %v, pc = %v\n", skip, pc[j])
        fmt.Printf("    file = %v, line = %d\n", file, line)
        fmt.Printf("    entry = %v\n", p.Entry())
        fmt.Printf("    name = %v\n", p.Name())

     }
        // skip = 0, pc = [4220961 4709117 4707856 4356842 4510849]
       
        // skip = 0, pc = 4220961
        //     file = /usr/local/go/src/runtime/extern.go, line = 217
        //     entry = 4220864
        //     name = runtime.Callers
        // skip = 0, pc = 4709117
        //     file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/7.runtime.go, line = 22
        //     entry = 4707872
        //     name = main.fun1
        // skip = 0, pc = 4707856
        //     file = /data/golang/src/github.com/xingcuntian/go_test/go-example/runtime/7.runtime.go, line = 18
        //     entry = 4707824
        //     name = main.main
        // skip = 0, pc = 4356842
        //     file = /usr/local/go/src/runtime/proc.go, line = 106
        //     entry = 4356320
        //     name = runtime.main
        // skip = 0, pc = 4510849
        //     file = /usr/local/go/src/runtime/asm_amd64.s, line = 2197
        //     entry = 4510848
        //     name = runtime.goexit

     break
 }   

}
package main

import (
    "fmt"
    "runtime"
)

//runtime.Callers 的用法
// func Callers(skip int, pc []uintptr) int
// 函数把当前go程调用栈上的调用栈标识符填入切片pc中，返回写入到pc中的项数。实参skip为开始在pc中记录之前所要跳过的栈帧数，
// 若为0则表示 runtime.Callers 自身的栈帧, 若为1则表示调用者的栈帧. 该函数返回写入到 pc 切片中的项数(受切片的容量限制).

func main(){
    fun1()
}

func fun1(){
    pc := make([]uintptr, 1024)
    for skip :=0; ; skip++ {
        n := runtime.Callers(skip,pc)
        if n <= 0{
            break
        }
      fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
    }
}
// skip = 0, pc = [4220961 4706990 4706832 4356842 4510737]
// skip = 1, pc = [4706990 4706832 4356842 4510737]
// skip = 2, pc = [4706832 4356842 4510737]
// skip = 3, pc = [4356842 4510737]
// skip = 4, pc = [4510737]
package main

import (
	"fmt"
	"sync/atomic"
)

// 交换Swap
// 与CAS操作不同，原子交换操作不会关心被操作的旧值。
// 它会直接设置新值
// 它会返回被操作值的旧值
// 此类操作比CAS操作的约束更少，同时又比原子载入操作的功能更强
var value int32

func main() {
	fmt.Println("=======old=======")
	fmt.Println(value)
	old := atomic.SwapInt32(&value, 10)
	fmt.Println("=========return old=========")
	fmt.Println(old)
	fmt.Println("======new value====")
	fmt.Println(value)
}

// =======old=======
// 0
// =========return old=========
// 0
// ======new value====
// 10

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int64 = 5
	var pn = &n
	var pf = (*float64)(unsafe.Pointer(pn))
	//现在，pn和pf指向同一个内存地址
    fmt.Println(pf) //0xc42000e268
	fmt.Println(pn) //0xc42000e268

	fmt.Println(*pf) //2.5e-323
	fmt.Println(*pn) //5
	*pf = 3.14159
	fmt.Println(n) //4614256650576692846
}
// 在这个例子中的转换可能是无意义的，但它是安全和合法的（为什么它是安全的？）。
// 因此，资源在unsafe包中的作用是为Go编译器服务，unsafe.Pointer类型的作用是绕过Go类型系统和内存安全。

	


package main

import (
	"fmt"
	"unsafe"
)

//unsafe.Pointer是一种特殊意义的指针，它可以包含任意类型的地址，有点类似于C语言里的void*指针，全能型的。
// 任何指针都可以转换为unsafe.Pointer
// unsafe.Pointer可以转换为任何指针
// uintptr可以转换为unsafe.Pointer
// unsafe.Pointer可以转换为uintptr
func main() {
	i := 10
	ip := &i

	var fp *float64 = (*float64)(unsafe.Pointer(ip))

	fmt.Println(*fp)

	*fp = *fp * 3

	fmt.Println(i)

}

// 以上示例，我们可以把*int转为*float64,并且我们尝试了对新的*float64进行操作，打印输出i，就会发现i的址同样被改变。

// 以上这个例子没有任何实际的意义，但是我们说明了，通过unsafe.Pointer这个万能的指针，我们可以在*T之间做任何转换。

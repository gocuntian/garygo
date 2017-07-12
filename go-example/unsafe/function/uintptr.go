package main

import (
	"fmt"
	"unsafe"
)

// uintptr可以转换为unsafe.Pointer
// unsafe.Pointer可以转换为uintptr
type user struct {
	name string
	age  int
	len  int
}

func main() {
	u := new(user)
	fmt.Println(*u) //{ 0}

	pName := (*string)(unsafe.Pointer(u))
	*pName = "星存田"
	fmt.Println(*u)

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	pLen := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.len)))
	*pLen = 100
	fmt.Println(*u)

}

//     temp:=uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)
//     pAge:=(*int)(unsafe.Pointer(temp))
//     *pAge = 20
// 逻辑上看，以上代码不会有什么问题，但是这里会牵涉到GC，如果我们的这些临时变量被GC，那么导致的内存操作就错了，我们最终操作的，就不知道是哪块内存了，会引起莫名其妙的问题。

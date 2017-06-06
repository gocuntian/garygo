package main

import (
	"fmt"
	"unsafe"
)
//由于uintptr是一个整数类型，uintptr值可以进行算术运算。 
//所以通过使用uintptr和unsafe.Pointer，我们可以绕过限制，* T值不能在Golang中计算偏移量：
func main(){
	a := []int{0,1,2,3}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2 * unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a = ",a)//a =  [0 1 2 6]

	type Person struct {
		name string
		age	int
		gender bool
	}
	who:= Person{"John",30,true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))

	*pname = "Alice"
	*page = 29
	*pgender = false
	fmt.Println(who)//{Alice 29 false}
}
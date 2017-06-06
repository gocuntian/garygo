package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	//将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	//保存结构体的地址备用(偏移量为 0)
	up0 := uintptr(p)
	//将通用指针转换为byte类型指针
	pb := (*byte)(p)
	//给转换后的指针赋值
	*pb = 10
	//结构体内容跟着改变
	fmt.Println(s) //{10 0 0 0}

	//偏移到第2个字段
	up := up0 + unsafe.Offsetof(s.b)
	//将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 类型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	fmt.Println(s) //{10 20 0 0}

	//偏移到第三个字段
	up = up0 + unsafe.Offsetof(s.c)
	//将偏移后的地址转换为通用地址
	p = unsafe.Pointer(up)
	//将通用地址转换为 byte 类型指针
	pb = (*byte)(p)
	//给转换后的指针赋值
	*pb = 30
	fmt.Println(s) //{10 20 30 0}

	up = up0 + unsafe.Offsetof(s.d)
	p = unsafe.Pointer(up)
	pi := (*int64)(p)
	*pi = 40
	fmt.Println(s) //{10 20 30 40}

}

// 结构体成员的内存分配是连续的，第一个成员的地址就是结构体的地址，相对于结构体的偏移量为 0。其它成员都可以通过偏移量来计算其地址。
// 每种类型都有它的大小和对齐值，可以通过 unsafe.Sizeof 获取其大小，通过 unsafe.Alignof 获取其对齐值，
// 通过 unsafe.Offsetof 获取其偏移量。不过 unsafe.Alignof 获取到的对齐值只是该类型单独使用时的对齐值，
// 不是作为结构体字段时与其它对象间的对齐值，这里用不上，所以需要用 unsafe.Offsetof 来获取字段的偏移量，进而确定其内存地址。

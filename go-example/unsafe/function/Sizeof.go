package main

import (
	"fmt"
	"unsafe"
)
//Sizeof函数可以返回一个类型所占用的内存大小，这个大小只有类型有关，
//和类型对应的变量存储的内容大小无关，比如bool型占用一个字节、int8也占用一个字节。

func main() {
	fmt.Println(unsafe.Sizeof(true)) // 1 byte =  8 bit
	fmt.Println(unsafe.Sizeof(int8(0))) //   1 byte =  8 bit
	fmt.Println(unsafe.Sizeof(int16(10))) // 2 byte = 16 bit
	fmt.Println(unsafe.Sizeof(int32(100000000))) // 4 byte = 32bit
	fmt.Println(unsafe.Sizeof(int64(1000000000000))) //8 byte = 64bit
	fmt.Println(unsafe.Sizeof(int(1000000000000000))) //64 位   8 byte = 64 bit
}
package main

/*
#include <stdlib.h>
char *cstring = "C string example";
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//2) Golang的字符串转C语言字符串
//通过C.Cstring和C.GoString来相互转化
func main() {
	var gstring string = "Go string example"
	cs := C.CString(gstring)         //Go string to C string
	defer C.free(unsafe.Pointer(cs)) // 需要手工释放cgo申请的字符串内存
	fmt.Println(cs)                  // 0x19ab1a0
	fmt.Println(string(*cs))         //G

	gs := C.GoString(C.cstring) // C string to Go string
	fmt.Println(gs)             // C string example

	gs2 := C.GoStringN(C.cstring, (C.int)(len(gs)-5))
	fmt.Println(gs2) // C string ex
}

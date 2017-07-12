package main

import (
	"fmt"
	"unsafe"
)

//Offsetof函数只适用于struct结构体中的字段相对于结构体的内存位置偏移量。结构体的第一个字段的偏移量都是0.
//此外，unsafe.Offsetof(u1.i)等价于reflect.TypeOf(u1).Field(i).Offset
type user1 struct {
	b byte  // 0
	i int32 // 4
	j int64 // 8
}

type user2 struct {
	b byte
	j int64
	i int32
}
type user3 struct {
	i int32
	b byte
	j int64
}
type user4 struct {
	i int32
	j int64
	b byte
}
type user5 struct {
	j int64
	b byte
	i int32
}
type user6 struct {
	j int64
	i int32
	b byte
}

func main() {
	var u1 user1
	fmt.Println("==============================1=============================")
	fmt.Println(unsafe.Offsetof(u1.b)) // 0
	fmt.Println(unsafe.Offsetof(u1.i)) // 4
	fmt.Println(unsafe.Offsetof(u1.j)) // 8
	// fmt.Println(reflect.TypeOf(u1).Field(0).Offset) // 0
	// fmt.Println(reflect.TypeOf(u1).Field(1).Offset) // 4
	// fmt.Println(reflect.TypeOf(u1).Field(2).Offset) // 8

	fmt.Println("=============================// ------ bxxx|iiii|jjjj|jjjj  ----------//============================")
	var u2 user2
	fmt.Println("==============================2=============================")
	fmt.Println(unsafe.Offsetof(u2.b)) // 0
	fmt.Println(unsafe.Offsetof(u2.j)) // 8
	fmt.Println(unsafe.Offsetof(u2.i)) // 16
	fmt.Println("==============================// ------ bxxx|xxxx|jjjj|jjjj|iiii|xxxx  ----------//=============================")
	fmt.Println("u1 size is ", unsafe.Sizeof(u1)) //16
	fmt.Println("u2 size is ", unsafe.Sizeof(u2)) // 24

	var u3 user3
	fmt.Println("==============================3=============================")
	fmt.Println("u3 size is ", unsafe.Sizeof(u3))
	fmt.Println(unsafe.Offsetof(u3.i)) // 0
	fmt.Println(unsafe.Offsetof(u3.b)) // 4
	fmt.Println(unsafe.Offsetof(u3.j)) // 8
	fmt.Println("==============================// ------ iiii|bxxx|jjjj|jjjj  ----------//=============================")
	var u4 user4
	fmt.Println("u4 size is ", unsafe.Sizeof(u4))
	fmt.Println(unsafe.Offsetof(u4.i)) // 0
	fmt.Println(unsafe.Offsetof(u4.j)) // 8
	fmt.Println(unsafe.Offsetof(u4.b)) // 16
	fmt.Println("==============================// ------iiii|xxxx|jjjj|jjjj|bxxx|xxxx  ----------//=============================")
	var u5 user5
	fmt.Println("u5 size is ", unsafe.Sizeof(u5))
	fmt.Println("==============================// ------jjjj|jjjj|bxxx|iiii  ----------//=============================")
	var u6 user6
	fmt.Println("u6 size is ", unsafe.Sizeof(u6))
	fmt.Println("==============================// ------jjjj|jjjj|iiii|bxxx  ----------//=============================")

}

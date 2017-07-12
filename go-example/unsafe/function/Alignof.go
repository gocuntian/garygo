package main

import (
	"fmt"
	"unsafe"
	"reflect"
)
//Alignof返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数。
//对齐值是一个和内存对齐有关的值，合理的内存对齐可以提高内存读写的性能\
//此外，获取对齐值还可以使用反射包的函数，也就是说：unsafe.Alignof(x)等价于reflect.TypeOf(x).Align()。
func main() {
	var b bool

	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32
	var f64 float64

	var s string

	var m map[string]string
	var p *int32
	var p64 *int64

	fmt.Println(unsafe.Alignof(b)) // 1

	fmt.Println(unsafe.Alignof(i8)) // 1
	fmt.Println(unsafe.Alignof(i16)) // 2
	fmt.Println(unsafe.Alignof(i64)) // 8

	fmt.Println(unsafe.Alignof(f32)) // 4
	fmt.Println(unsafe.Alignof(f64)) // 8

	fmt.Println(unsafe.Alignof(s)) // 8

	fmt.Println(unsafe.Alignof(m)) // 8

	fmt.Println(unsafe.Alignof(p)) // 8
	fmt.Println(unsafe.Alignof(p64)) // 8

	fmt.Println(reflect.TypeOf(p).Align()) //8

}
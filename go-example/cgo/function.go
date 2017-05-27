// 5) Golang访问C语言的function
// Golang通过cgo可以直接访问C语言定义的函数, 其中传入的参数也必须转成C语言的类型，
// 不能使用Golang的原生类型。同时直接访问也不支持C语言的union，宏等特性，
// 需要写一个wrapper来屏蔽接口，C++语言访问也经常通过C语言来中转。
package main

/*
#include <stdio.h>
#include <stdint.h>
struct columns {
	int column1;
	int column2;
	int column3;
};
int sum_columns(struct columns a) {
	return a.column1 + a.column2 + a.column3;
}
int sum_vals(int a, int b) {
	return a + b;
}
*/
import "C"
import "fmt"

func main() {
	c := C.struct_columns{15, 30, 45}
	sum := C.sum_columns(c)
	fmt.Println(sum) //90
	var a int = 15
	var b int = 30
	s := C.sum_vals((C.int)(a), (C.int)(b)) // 调用C函数
	fmt.Println(s)                          //45
	var goSum int = int(sum)                // C int to Go int
	fmt.Println(goSum)                      //90
}

// 为了优化GC,可以借助unsafe.Pointer在cgo内部申请对象，释放对象，Golang只需调用相应方法。

// var dst unsafe.Pointer
// ret := C.call(&dst) //通常在C语言内部，申请C++对象，绑定到dst上，返回给
// //balala...
// //C.real_call(dst, arg) //内部调用dst->func(arg)方法
// //balala...
// C.free(dst)
// 使用cgo的考虑

// 1）Golang runtime不会处理cgo模块涉及的内存，适当使用cgo可以减少gc压力，降低Golang延时。
// 2）Cgo的使用也不是没有代价的，通常表明一次cgo调用耗时170ns，而常规Golang调用只需9ns。在gortoutine内部调用cgo的sleep()会同步等待，影响并发。同时会失去Golang跨平台的能力。

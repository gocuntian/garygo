package main

/*
#include <stdlib.h>
struct Person {
	char *name;
	int	age;
	int height;
	int weight;
};
*/
import "C"
import "fmt"

type p C.struct_Person

//3)Golang访问C语言的struct
// C语言的struct在Golang中可以直接访问，对应C.struct_xxxx，其中xxxx是C语言的struct名称。
// 注意一下struct的内存对齐问题。

func main() {
	person := p{C.CString("xingcuntian"), 30, 6, 170, [4]byte{}}
	//最后4byte是自动转化时内存对齐造成的.
	fmt.Println(person)
	fmt.Println(C.GoString(person.name))
	fmt.Println(person.age)
	fmt.Println(person.height)
	fmt.Println(person.weight)
}

//{0x1e68050 30 6 170 [0 0 0 0]}
// xingcuntian
// 30
// 6
// 170

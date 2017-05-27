package main

/*
enum levels {
	low,
	medium,
	high
};

typedef enum {
	LOW = 0,
	MEDIUM = 1,
	HIGH = 2
} secutiry;
*/
import "C"

// 4)Golang访问C语言的enum
// Golang语言没有原生的enum支持，只能通过const来模拟。然后通过cgo我们可以使用enum。
import "fmt"

func main() {
	cc := new(C.enum_levels)
	fmt.Println(*cc)      // 0
	fmt.Println(C.LOW)    //0
	fmt.Println(C.MEDIUM) // 1
	fmt.Println(C.HIGH)   // 2
}

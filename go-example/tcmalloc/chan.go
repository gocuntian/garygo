package main

import "fmt"

func main() {
	a := make([]int, 0, 20)
	fmt.Println("cap=", cap(a), "len=", len(a))
	b := make([]int, 0, 20000)
	fmt.Println("cap=", cap(b), "len=", len(b))
	l := 20
	c := make([]int, 0, l)
	fmt.Println("cap=", cap(c), "len=", len(c))
	d := make([]int, 0, 8191)
	fmt.Println("cap=", cap(d), "len=", len(d))
}

// go build -gcflags="-m" chan.go 2>&1
// a 和 b的代码一样，就是申请的空间不一样大， 但是它们两个的命运是截然相反的。
// a前面已经介绍过了， 会申请到栈上， 而b由于申请内存较大， 编译器会把这种申请内存较大（>8191）的变量转移到堆上面。
// 即使是临时变量， 申请过大也会在堆上面申请. 而c， 对我们而言， 其含义和a是一致的， 但是编译器对于这种不定长度的申请方式， 也会在堆上面申请， 即使申请的长度很短。
// 可以通过下面的命令查看变量申请的位置。 详细内容可以参考我之前的文章《【译】优化Go的模式》

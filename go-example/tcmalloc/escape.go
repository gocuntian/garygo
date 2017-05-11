package main

import ()

func foo() *int {
    var x int
    return &x
}

func bar() int {
    x := new(int)
    *x = 1
    return *x
}

func main(){}
//  go run -gcflags '-m -l' escape.go
// # command-line-arguments
// ./escape.go:7: &x escapes to heap
// ./escape.go:6: moved to heap: x
// ./escape.go:11: bar new(int) does not escape


// 如何得知变量是分配在栈（stack）上还是堆（heap）上？

// 准确地说，你并不需要知道。Golang 中的变量只要被引用就一直会存活，存储在堆上还是栈上由内部实现决定而和具体的语法没有关系。

// 知道变量的存储位置确实和效率编程有关系。如果可能，Golang 编译器会将函数的局部变量分配到函数栈帧（stack frame）上。然而，如果编译器不能确保变量在函数 return 之后不再被引用，编译器就会将变量分配到堆上。而且，如果一个局部变量非常大，那么它也应该被分配到堆上而不是栈上。

// 当前情况下，如果一个变量被取地址，那么它就有可能被分配到堆上。然而，还要对这些变量做逃逸分析，如果函数 return 之后，变量不再被引用，则将其分配到栈上。
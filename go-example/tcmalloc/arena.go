package main

func foo() *int{
    x := 1
    return &x
}

func main(){
    x := foo()
    println(*x)
}
// go build -gcflags '-l' -o arena arena.go
// go tool objdump -s "main\.foo" arena
//堆上内存分配调用了 runtime 包的 newobject 函数。
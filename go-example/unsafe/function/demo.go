package main

import (
	"fmt"
)

//同时为了安全的考虑，Go语言是不允许两个指针类型进行转换的。
func main() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(ip)
	fmt.Println(fp)
}

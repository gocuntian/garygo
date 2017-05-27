package main

import (
	"fmt"
)

func main() {
	a := make([]int64, 0)
	fmt.Println("cap=", cap(a), "len=", len(a))

	for i := 0; i < 3; i++ {
		a = append(a, 1)
		fmt.Println("cap=", cap(a), "len=", len(a))
	}
}

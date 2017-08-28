package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	for k, v := range x {
		fmt.Printf("x[%d] = %d\n", k, v)
	}
}

// x[0] = 10
// x[1] = 20
// x[2] = 30
// x[3] = 40
// x[4] = 50

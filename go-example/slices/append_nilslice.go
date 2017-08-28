package main

import (
	"fmt"
)

func main() {
	var x []int
	fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  0 cap =  0
	x = append(x, 10)
	fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  1 cap =  1

	// x = append(x, 3, 4)
	// fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  3 cap =  4

	// x = append(x, 20)
	// fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  2 cap =  2
	// x = append(x, 30, 40)
	// fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  4 cap =  4
	// x = append(x, 50, 60, 70)
	// fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  7 cap =  8
	// x = append(x, 80, 90, 100, 110, 120)
	// fmt.Println("Slice x len = ", len(x), "cap = ", cap(x)) //Slice x len =  12 cap =  16
}

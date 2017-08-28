package main

import (
	"fmt"
)

func main() {
	x := make([]int, 2, 5)
	fmt.Println("Slice x: len=", len(x), "cap=", cap(x)) //Slice x: len= 2 cap= 5
	x[0] = 10
	x[1] = 20
	fmt.Println("Slice x:", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x)) //Length is 2 Capacity is 5
	x = append(x, 30, 40, 50)                                   // Slice x [0 0 30 40 50]  //Slice x [10 20 30 40 50]
	fmt.Println("Slice x", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x)) //Length is 5 Capacity is 5

	x = append(x, 60)
	fmt.Println("Slice x", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x)) //Length is 6 Capacity is 10

	x = append(x, 70, 80, 90, 100)
	fmt.Println("Slice x", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x)) //Length is 10 Capacity is 10

	x = append(x, 110)
	fmt.Println("Slice x", x)
	fmt.Printf("Length is %d Capacity is %d\n", len(x), cap(x)) //Length is 11 Capacity is 20
}

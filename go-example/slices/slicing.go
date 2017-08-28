package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice x = ", x, ",len = ", len(x), ", cap = ", cap(x)) //Slice x =  [10 20 30 40 50] ,len =  5 , cap =  5
	y := x[1:3]
	fmt.Println("Slice y = ", y, ",len = ", len(y), ", cap = ", cap(y)) //Slice y =  [20 30] ,len =  2 , cap =  4
	z := x[:3]
	fmt.Println("Slice z = ", z, ",len = ", len(z), ", cap = ", cap(z)) //Slice z =  [10 20 30] ,len =  3 , cap =  5
	q := x[:4]
	fmt.Println("Slice q = ", q, ",len = ", len(q), ", cap = ", cap(q)) // Slice q =  [10 20 30 40] ,len =  4 , cap =  5
	x1 := x[:]
	fmt.Println("Slice x1 = ", x1, ",len = ", len(x1), ", cap = ", cap(x)) //Slice x1 =  [10 20 30 40 50] ,len =  5 , cap =  5

	x1[4] = 75
	fmt.Println("x:", x)   //x: [10 20 30 40 75]
	fmt.Println("x1:", x1) //x1: [10 20 30 40 75]
}

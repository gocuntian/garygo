package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30}
	fmt.Println("Slice x = ", x, ",len = ", len(x), ", cap = ", cap(x)) //Slice x =  [10 20 30] ,len =  3 , cap =  3
	y := make([]int, 5, 10)
	copy(y, x)
	fmt.Printf("[Slice:y] Length is %d Capacity is %d\n", len(y), cap(y)) //[Slice:y] Length is 5 Capacity is 10
	fmt.Println("Slice y after copying:", y)                              //Slice y after copying: [10 20 30 0 0]
	y[3] = 40
	y[4] = 50
	fmt.Println("Slice y after adding elements:", y) //Slice y after adding elements: [10 20 30 40 50 60]
	//y[5] = 60 //anic: runtime error: index out of range
	y = append(y, 60)
	fmt.Printf("[Slice:y] Length is %d Capacity is %d\n", len(y), cap(y)) //[Slice:y] Length is 6 Capacity is 10
	fmt.Println("Slice y after appending elements:", y)                   //Slice y after adding elements: [10 20 30 40 50]
}

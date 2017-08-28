package main

import "fmt"

func main() {
	z := []int{0: 10, 2: 20}
	fmt.Println("Slice z= ", z, ", len = ", len(z), ", cap = ", cap(z)) //Slice z=  [10 0 20] , len =  3 , cap =  3
}

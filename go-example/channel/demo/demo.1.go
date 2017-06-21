package main

import (
	"fmt"
)

func add(a,b int) <-chan int {
	sum := make(chan int)
	go func() {
		sum<-a+b
	}()
	return sum
}

func main() {
	sum12 := add(1,2)
	sum45 := add(4,5)
	fmt.Println(<-sum12)
	fmt.Println(<-sum45)
}
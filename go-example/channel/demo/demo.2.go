package main

import (
	"fmt"
)

func fanln(sum1, sum2 <-chan int) <-chan int {
	sum := make(chan int)
	go func() {
		for {sum<- <-sum1}
	}()
	go func() {
		for {sum<- <-sum2}
	}()
	return sum
}

func add(a,b int) <-chan int {
	sum := make(chan int)
	go func() {
		sum<-a+b
	}()
	return sum
}

func main() {
	sum:=fanln(add(1,2),add(4,5))
	fmt.Println(<-sum)
	fmt.Println(<-sum)
}
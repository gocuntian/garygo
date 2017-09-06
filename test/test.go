package main

import (
	"fmt"
	"strconv"
)

var a, b, c, d string
var num = [4]int{1, 2, 3, 4}
var strings = [4]string{a, b, c, d}

func main() {
	range_num(num)
	fmt.Println(num)
	for i := 0; i < 11; i++ {
		num[0], num[1], num[2], num[3] = num[1], num[2], num[3], num[0]
		fmt.Println(num)
		range_num(num)
	}
	fmt.Println(strings)
}

func write(i int, s *string) {
	*s += strconv.Itoa(i)
}

func range_num(num [4]int) {
	for i := 0; i < 4; i++ {
		fmt.Println(num[i])
		write(num[i], &strings[i])
		fmt.Println(strings)
	}
}

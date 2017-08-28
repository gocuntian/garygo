package main

import (
	"fmt"
)

func main() {
	var chapts map[int]string
	fmt.Println(chapts)
	chapts = make(map[int]string)
	fmt.Println(chapts)

	chapts[1] = "Beginning Go"
	chapts[2] = "Go Fundamentals"
	chapts[3] = "Structs and Interfaces"
	delete(chapts, 1)
	fmt.Println(chapts)
}

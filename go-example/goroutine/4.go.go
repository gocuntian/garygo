package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"one", "two", "three", "four", "five"}
	for _, name := range names {
		go func(name string) {
			fmt.Println("this is :", name)
		}(name)
	}
	runtime.Gosched()
}

// this is : one
// this is : two
// this is : four
// this is : three
// this is : five

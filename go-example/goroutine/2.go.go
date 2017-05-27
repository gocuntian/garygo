package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello, %s!\n", name)
	}()
	name = "Harry"
	time.Sleep(time.Millisecond)
}

//Hello, Harry!
// goroutine 的执行顺序是不可预知的

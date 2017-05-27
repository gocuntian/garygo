package main

import (
	"runtime"
	//"time"
)

func main() {
	go println("Go! Goroutine!\r\n")
	//time.Sleep(time.Millisecond)
	runtime.Gosched()
}

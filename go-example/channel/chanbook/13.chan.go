package main

import (
	"fmt"
	"time"
)

//超时
func main() {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)
		case <-timer.C:
			fmt.Println("Timeout!")
		}
	}
}

// Received: 0
// Timeout!
// Received: 1
// Timeout!
// Received: 2
// Timeout!
// Received: 3
// Timeout!
// Received: 4
// End.

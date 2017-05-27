package main

import (
	"fmt"
	"time"
)

//非缓冲通道
func main() {
	intChan := make(chan int)
	syncChan := make(chan bool)
	ticker := time.NewTicker(time.Second)
	go func() {
	Look:
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			case <-syncChan:
				ticker.Stop()
				close(intChan)
				break Look
			}
		}
		fmt.Println("End. [sender]")

	}()

	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			syncChan <- true
		}
	}
	fmt.Println("End. [receiver]")
}

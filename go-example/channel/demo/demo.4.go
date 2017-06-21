package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { wg.Done() }()
	go func() { wg.Done() }()
	wg.Wait()
}

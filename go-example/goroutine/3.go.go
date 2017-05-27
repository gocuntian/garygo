package main

import (
	"fmt"
	"runtime"
)

func main() {
	names := []string{"one", "two", "three", "four", "five"}
	for _, name := range names {
		go func() {
			fmt.Println("name:", name)
		}()
	}
	runtime.Gosched()
}

// name: five
// name: five
// name: five
// name: five
// name: five

// go run -race 3.go.go
// ==================
// WARNING: DATA RACE
// Read at 0x00c42007c190 by goroutine 6:
//   main.main.func1()
//       /data/go/src/github.com/xingcuntian/go_test/go-example/goroutine/3.go.go:12 +0x68

// Previous write at 0x00c42007c190 by main goroutine:
//   main.main()
//       /data/go/src/github.com/xingcuntian/go_test/go-example/goroutine/3.go.go:10 +0xf4

// Goroutine 6 (running) created at:
//   main.main()
//       /data/go/src/github.com/xingcuntian/go_test/go-example/goroutine/3.go.go:13 +0x131
// ==================
// name: two
// name: three
// name: five
// name: five
// name: five
// Found 1 data race(s)
// exit status 66

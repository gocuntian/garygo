package main

import (
	"fmt"

	"time"

	"github.com/silentred/gid"
)

func main() {
	for i := 0; i < 10; i++ {
		go test(i)
	}
	time.Sleep(3 * time.Second)
	// id := gid.Get()
	//fmt.Println(id)
}

func test(i int) {
	id := gid.Get()
	fmt.Println("code: ", i, "gid: ", id)
}

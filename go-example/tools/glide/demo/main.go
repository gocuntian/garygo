package main

import (
	"fmt"

	"github.com/funny/slab"
)

func main() {
	pool := slab.NewSyncPool(
		64,      // The smallest chunk size is 64B.
		64*1024, // The largest chunk size is 64KB.
		2,       // Power of 2 growth in chunk size.
	)
	fmt.Println(pool)
	buf := pool.Alloc(64)

	pool.Free(buf)
	fmt.Println("ok")
}

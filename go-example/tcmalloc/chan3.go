package main

import (
	"fmt"
	"sync"
)

var idsPool = sync.Pool{
	New: func() interface{} {
		ids := make([]int64, 0, 20000)
		return &ids
	},
}

func NewIds() []int64 {
	ids := idsPool.Get().(*[]int64)
	*ids = (*ids)[:0]
	idsPool.Put(ids)
	return *ids
}

func main() {
	b := make([]int64, 0, 20000)
	fmt.Println("cap=", cap(b), "len=", len(b))
	c := NewIds()
	fmt.Println("cap=", cap(c), "len=", len(c))
}

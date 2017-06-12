package slab

import (
	"fmt"
	"unsafe"
)

type ChanPool struct {
	classes []chanClass
	minSize int
	maxSize int
}

func NewChanPool(minSize, maxSize, factor, pageSize int) *ChanPool {
	pool := &ChanPool{make([]chanClass, 0, 10), minSize, maxSize}
	for chunkSize := minSize; chunkSize <= maxSize && chunkSize <= pageSize; chunkSize *= factor {
		c := chanClass{
			size:   chunkSize,
			page:   make([]byte, pageSize),
			chunks: make(chan []byte, pageSize/chunkSize),
		}
		c.pageBegin = uintptr(unsafe.Pointer(&c.page[0]))
		for i := 0; i < pageSize/chunkSize; i++ {
			mem := c.page[i*chunkSize : (i+1)*chunkSize : (i+1)*chunkSize]
			fmt.Println(mem)
			c.chunks <- mem
			if i == len(c.chunks)-1 {
				c.pageEnd = uintptr(unsafe.Pointer(&mem[0]))
			}
		}
		pool.classes = append(pool.classes, c)
	}
	return pool
}

func (pool *ChanPool) Alloc(size int) []byte {
	if size <= pool.maxSize {
		for i := 0; i < len(pool.classes); i++ {
			if pool.classes[i].size >= size {
				mem := pool.classes[i].Pop()
				if mem != nil {
					return mem[:size]
				}
				break
			}
		}
	}
	return make([]byte, size)
}

func (pool *ChanPool) Free(mem []byte) {
	size := cap(mem)
	for i := 0; i < len(pool.classes); i++ {
		if pool.classes[i].size == size {
			pool.classes[i].Push(mem)
			break
		}
	}

}

type chanClass struct {
	size      int
	page      []byte
	pageBegin uintptr
	pageEnd   uintptr
	chunks    chan []byte
}

func (c *chanClass) Push(mem []byte) {
	c.chunks <- mem
}

func (c *chanClass) Pop() []byte {
	select {
	case mem := <-c.chunks:
		return mem
	default:
		return nil
	}
}

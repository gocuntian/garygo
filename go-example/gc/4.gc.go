package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	_ = c
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	runtime.GC()
	memstats := new(runtime.MemStats)
	runtime.ReadMemStats(memstats)
	fmt.Println(memstats.Alloc)
	fmt.Println(memstats.TotalAlloc)
	fmt.Println(memstats.Mallocs)
	fmt.Println(memstats.HeapAlloc)
	fmt.Println(memstats.HeapObjects)
	fmt.Println(memstats.HeapSys)
	fmt.Println(memstats.HeapIdle)

}

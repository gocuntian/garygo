package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	const N = 10000
	st := new(runtime.MemStats)
	memstats := new(runtime.MemStats)
	fmt.Println(memstats)
	runtime.ReadMemStats(st)
	fmt.Println(memstats.Alloc)
	fmt.Println(memstats.TotalAlloc)
	fmt.Println(memstats.Mallocs)
	fmt.Println(memstats.HeapAlloc)
	fmt.Println(memstats.HeapObjects)
	fmt.Println(memstats.HeapSys)
	fmt.Println(memstats.HeapIdle)
	for i := 0; i < N; i++ {
		c := make(chan int, 10)
		_ = c
		if i%100 == 0 {
			for j := 0; i < 4; j++ {
				runtime.GC()
				runtime.Gosched()
				runtime.GC()
				runtime.Gosched()
			}
		}
	}

	fmt.Println(memstats)

	runtime.ReadMemStats(memstats)
	obj := int64(memstats.HeapObjects - st.HeapObjects)
	if obj > N/5 {
		fmt.Println("too many objects left:", obj)
		os.Exit(1)
	}

}

// type MemStats struct {
//     // 一般统计
//     Alloc      uint64 // 已申请且仍在使用的字节数
//     TotalAlloc uint64 // 已申请的总字节数（已释放的部分也算在内）
//     Sys        uint64 // 从系统中获取的字节数（下面XxxSys之和）
//     Lookups    uint64 // 指针查找的次数
//     Mallocs    uint64 // 申请内存的次数
//     Frees      uint64 // 释放内存的次数
//     // 主分配堆统计
//     HeapAlloc    uint64 // 已申请且仍在使用的字节数
//     HeapSys      uint64 // 从系统中获取的字节数
//     HeapIdle     uint64 // 闲置span中的字节数
//     HeapInuse    uint64 // 非闲置span中的字节数
//     HeapReleased uint64 // 释放到系统的字节数
//     HeapObjects  uint64 // 已分配对象的总个数
//     // L低层次、大小固定的结构体分配器统计，Inuse为正在使用的字节数，Sys为从系统获取的字节数
//     StackInuse  uint64 // 引导程序的堆栈
//     StackSys    uint64
//     MSpanInuse  uint64 // mspan结构体
//     MSpanSys    uint64
//     MCacheInuse uint64 // mcache结构体
//     MCacheSys   uint64
//     BuckHashSys uint64 // profile桶散列表
//     GCSys       uint64 // GC元数据
//     OtherSys    uint64 // 其他系统申请
//     // 垃圾收集器统计
//     NextGC       uint64 // 会在HeapAlloc字段到达该值（字节数）时运行下次GC
//     LastGC       uint64 // 上次运行的绝对时间（纳秒）
//     PauseTotalNs uint64
//     PauseNs      [256]uint64 // 近期GC暂停时间的循环缓冲，最近一次在[(NumGC+255)%256]
//     NumGC        uint32
//     EnableGC     bool
//     DebugGC      bool
//     // 每次申请的字节数的统计，61是C代码中的尺寸分级数
//     BySize [61]struct {
//         Size    uint32
//         Mallocs uint64
//         Frees   uint64
//     }
// }

// func ReadMemStats(m *MemStats)
// ReadMemStats将内存申请和分配的统计信息填写进m。

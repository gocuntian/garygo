package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var data *string

func Data() string {
	p := (*string)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if p == nil {
		return ""
	} else {
		return *p
	}
}

func SetData(d string) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
		unsafe.Pointer(&d),
	)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)

	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			log.Println(Data())
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			log.Println("===", s)
			SetData(s)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("final data = ", *data)
}

// 合法用例2: 调用sync/atomic包中指针相关的函数
// sync / atomic包中的以下函数的大多数参数和结果类型都是unsafe.Pointer或*unsafe.Pointer：
// func CompareAndSwapPointer（addr * unsafe.Pointer，old，new unsafe.Pointer）（swapped bool）
// func LoadPointer（addr * unsafe.Pointer）（val unsafe.Pointer）
// func StorePointer（addr * unsafe.Pointer，val unsafe.Pointer）
// func SwapPointer（addr * unsafe.Pointer，new unsafe.Pointer）（old unsafe.Pointer）
// 要使用这些功能，必须导入unsafe包。
// 注意： unsafe.Pointer是一般类型，因此 unsafe.Pointer的值可以转换为unsafe.Pointer，反之亦然。

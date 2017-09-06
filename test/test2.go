package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//count:输出次数；name: 当前线程(协程)名称；source：来源chan，某文件句柄从该通道获取；目标chan：将某文件句柄push进该通道
func task(count int, name uint8, source chan []uint8, target chan []uint8) {
	for i := 0; i < count; i++ {
		find := <-source
		find = append(find, name)
		target <- find
	}
	wg.Done()
}

func main() {
	arrLen := 10 //循环次数定义为10
	//定义4个通道，用于4个线程(协程)间传递信息
	c1, c2, c3, c4 := make(chan []uint8, 1), make(chan []uint8, 1), make(chan []uint8, 1), make(chan []uint8, 1)
	//定义4个切片，模拟4个文件
	f1, f2, f3, f4 := make([]uint8, 0, arrLen), make([]uint8, 0, arrLen), make([]uint8, 0, arrLen), make([]uint8, 0, arrLen)
	//将4个切片push进通道
	c1 <- f1
	c2 <- f2
	c3 <- f3
	c4 <- f4
	//开始循环，启动4个线程(协程)
	arr := []chan []uint8{c1, c2, c3, c4}
	for i := 0; i < len(arr); i++ {
		if i == 3 {
			wg.Add(1)
			go task(10, uint8(i+1), arr[i], arr[0])
		} else {
			wg.Add(1)
			go task(10, uint8(i+1), arr[i], arr[i+1])
		}
	}
	fmt.Println("begin^_^")
	wg.Wait()
	//打印结果
	fmt.Println("f1: ", <-c1)
	fmt.Println("f2: ", <-c2)
	fmt.Println("f3: ", <-c3)
	fmt.Println("f4: ", <-c4)
	fmt.Println("exit^_^")

}

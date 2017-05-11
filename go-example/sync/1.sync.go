package main

import (
	"fmt"
	"sync"
)

// type WaitGroup struct {
//     // 包含隐藏或非导出字段
// }
// WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
// 每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。

// func (wg *WaitGroup) Add(delta int) //添加或者减少等待goroutine的数量
// Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，
// Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。
// 注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。
// 一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。

// func (wg *WaitGroup) Done() //相当于Add(-1)
// Done方法减少WaitGroup计数器的值，应在线程的最后执行。

// func (wg *WaitGroup) Wait() //执行阻塞，直到所有的WaitGroup数量变成0
// Wait方法阻塞直到WaitGroup计数器减为0。

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println(shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
	for i := 0; i < 10; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}

// 使用场景：
// 　　 程序中需要并发，需要创建多个goroutine，并且一定要等这些并发全部完成后才继续接下来的程序执行．
//     WaitGroup的特点是Wait()可以用来阻塞直到队列中的所有任务都完成时才解除阻塞，
//     而不需要sleep一个固定的时间来等待．但是其缺点是无法指定固定的goroutine数目

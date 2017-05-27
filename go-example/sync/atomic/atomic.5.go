package main

import (
	"fmt"
	"sync/atomic"
)
// 载入Load
// 上面的比较并交换案例总 v:= value为变量v赋值，但… 要注意，在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
// so, 我们要使用sync/atomic代码包同样为我们提供了一系列的函数，以Load为前缀(载入)，来确保这样的糟糕事情发

//载入 Loadxxxx

var value int32

func main(){
	fmt.Println("======old value=======")
    fmt.Println(value)
    fmt.Println("======CAS value=======")
    addValue(3)
    fmt.Println(value)
}

//不断地尝试原子地更新value的值,直到操作成功为止
func addValue(delta int32){
 //在被操作值被频繁变更的情况下,CAS操作并不那么容易成功
    //so 不得不利用for循环以进行多次尝试
	for {
		 //v := value
        //在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
        //因此我们要使用载
		v :=atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value,v,(v + delta)) {
			//在函数的结果值为true时,退出循环
			break
		}
		//操作失败的缘由总会是value的旧值已不与v的值相等了.
        //CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
	}
}
// func LoadInt32(addr *int32) (val int32)
// LoadInt32原子性的获取*addr的值。
// atomic.LoadInt32接受一个*int32类型的指针值
// 返回该指针指向的那个值

// ======old value=======
// 0
// ======CAS value=======
// 3

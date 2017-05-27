package main

import (
	"fmt"
	"sync/atomic"
)

// 存储Store
// 与读取操作相对应的是写入操作。 而sync/atomic包也提供了与原子的载入函数相对应的原子的值存储函数。 以Store为前缀
// 在原子的存储某个值的过程中，任何CPU都不会进行针对同一个值的读或写操作。
// 原子的值存储操作总会成功，因为它并不会关心被操作值的旧值是什么
// 和CAS操作有着明显的区别

var value int32
func main(){
	fmt.Println("=======Store============")
	atomic.StoreInt32(&value,10)
	fmt.Println(value)
}
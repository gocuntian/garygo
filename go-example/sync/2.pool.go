package main

import (
    "fmt"
    "runtime"
    "runtime/debug"
    "sync"
    "sync/atomic"
)

func main(){
    // 禁用GC，并保证在main函数执行结束前恢复GC
    defer debug.SetGCPercent(debug.SetGCPercent(-1))

    var count int32

    newFunc := func()interface{}{
        return atomic.AddInt32(&count,1)
    }
    pool := sync.Pool{New: newFunc}

    // New 字段值的作用
    // v1 := pool.Get()
    // fmt.Printf("v1:%v\n",v1)

    //临时对象池的存取
    pool.Put(newFunc())
    v2 := pool.Get()
    fmt.Printf("v2: %v\n", v2)
    pool.Put(newFunc())
    pool.Put(newFunc())
    // v3 := pool.Get()
    // fmt.Printf("v3: %v\n", v3)

    //垃圾回收对临时对象池的影响
    debug.SetGCPercent(100)
    runtime.GC()
    v4 := pool.Get()
    fmt.Printf("v4: %v\n", v4)
    pool.New = nil
    v5 := pool.Get()
    fmt.Printf("v5: %v\n", v5)

}

// AddT 系列函数实现加法操作，在原子性上等价于：
// *addr += delta
// return *addr

// func AddInt32(addr *int32, delta int32) (new int32)
// AddInt32原子性的将val的值添加到*addr并返回新值。


// 通过Get方法获取到的值是任意的。如果一个临时对象池的Put方法未被调用过，且它的New字段也未曾被赋予一个非nil的函数值，
// 那么它的Get方法返回的结果值就一定会是nil。Get方法返回的不一定就是存在于池中的值。不过，如果这个结果值是池中的，那么在该方法返回它之前就一定会把它从池中删除掉。
// 这样一个临时对象池在功能上看似与一个通用的缓存池相差无几。
// 但是实际上，临时对象池本身的特性决定了它是一个“个性”非常鲜明的同步工具。我们在这里说明它的两个非常突出的特性
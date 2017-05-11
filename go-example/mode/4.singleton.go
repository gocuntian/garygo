package main

import (
    "sync"
    "sync/atomic"
)
//因为编译器优化没有检查实例存储状态。如果使用sync/atomic包的话 就可以自动帮我们加载和设置标记。
var mx sync.Mutex
type singleton struct{}

var instance *singleton
var initialized uint32

func GetInstance() *singleton{
    
    if atomic.LoadUInt32(&initialized) == 1 {
        return instance
    }

    mx.Lock()
    defer mx.Unlock()

    if initialized == 0 {
        instance = &singleton{}
        atomic.StoreUint32(&initialized, 1)
    }
    return instance

}

// func LoadUint32(addr *uint32) (val uint32)
// LoadUint32原子性的获取*addr的值。
// func StoreUint32(addr *uint32, val uint32)
// StoreUint32原子性的将val的值保存到*addr。

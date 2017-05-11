package main

import (
    "sync"
)

//2.Go实现带线程锁的单例模式
type singleton struct{}
var mu sync.Mutex

var instance *singleton

func GetInstace() * singleton {
    mu.Lock()
    defer mu.Unlock()
    if instance == nil {
        instance = & singleton{}
    }
    return instance
}

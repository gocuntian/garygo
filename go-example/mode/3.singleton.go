package main

//3.带检查锁的的单例模式
import (
    "sync"
)
var mx sync.Mutex
type singleton struct{}

var instance *singleton

func GetInstace() * singleton{
    if instance == nil {
        mx.Lock()
        defer mx.Unlock()
        instance = & singleton{}
    }
    return instance
}



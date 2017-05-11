package main

//1.Go实现非线程安全的单例模式

type singleton struct{}

var instance *singleton

func GetInstace() *singleton{
    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
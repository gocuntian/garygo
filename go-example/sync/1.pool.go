package main

import (
    "log"
    "sync"
)
// var pool = &sync.Pool{New:func()interface{}{return NewObject()}}
//     pool.Put()
//     Pool.Get()
// 对象池在Get的时候没有里面没有对象会返回nil，所以我们需要New function来确保当获取对象对象池为空时，重新生成一个对象返回，
// 前者的功能是从池中获取一个interface{}类型的值，而后者的作用则是把一个interface{}类型的值放置于池中。
func main(){
    //创建对象
    var pool = &sync.Pool{New:func() interface{}{return "Hello,xingcuntian"}}
    //默认值Hello,xingcuntian
    log.Println(pool.Get())
    //准备放入的字符串
    val := "Hello, World!"
    //放入
    pool.Put(val)
    //取出
    log.Println(pool.Get())
    //再取就没有了,会自动调用NEW
    log.Println(pool.Get())
    log.Println(pool.Get())

}
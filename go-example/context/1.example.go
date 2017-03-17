package main

import (
    "fmt"
    "context"
)
//Example(WithCancel)
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc){}
// WithCancel返回具有新的完成通道的父级的副本。 返回的上下文的完成通道在调用返回的cancel函数时或当父上下文的完成通道关闭时（以先发生者为准）关闭。
// 取消此上下文释放与其相关联的资源，因此代码应在此上下文中运行的操作完成后立即调用cancel。

// gen在一个单独的goroutine中生成整数
//将它们发送到返回的通道。
// gen的调用者需要取消上下文一次
//他们做消耗生成的整数不泄漏
//由gen开始的内部goroutine。

func main(){
    gen := func(ctx context.Context) <-chan int {
        dst := make(chan int)
        n := 1
        go func(){
            for {
                select {
                    case <-ctx.Done():
                         return //返回不泄漏goroutine
                    case dst <-n:
                        n++
                }
            }
        }()
        return dst
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
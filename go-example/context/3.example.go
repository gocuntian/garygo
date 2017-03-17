package main

import (
    "fmt"
    "time"
    "context"
)
//Example(WithTimeout)
// 调用WithCancel()可以将基础的 Context 进行继承，返回一个cancelCtx示例，并返回一个函数，
// 可以在外层直接调用cancelCtx.cancel()来取消 Context
// 代码如下：
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
//     c := newCancelCtx(parent)
//     propagateCancel(parent, &c)
//     return &c, func() { c.cancel(true, Canceled) }
// }

func main(){
    ctx, cancel := context.WithTimeout(context.Background(),50 * time.Millisecond)
    defer cancel()

    select{
        case <-time.After(1 * time.Second):
             fmt.Println("overslept")
        case <-ctx.Done():
             fmt.Println(ctx)
             fmt.Println(<-ctx.Done())
             fmt.Println(ctx.Err())
    }
}

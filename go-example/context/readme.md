### 概述 Golang 的 context Package 提供了一种简洁又强大方式来管理 goroutine 的生命周期，同时提供了一种 Requst-Scope K-V Store。但是对于新手来说，Context 的概念不算非常的直观，这篇文章来带领大家了解一下 Context 包的基本作用和使用方法。 
### 1. 包的引入 在 go1.7 及以上版本 context 包被正式
#概述
 Golang 的 context Package 提供了一种简洁又强大方式来管理 goroutine 的生命周期，同时提供了一种 Requst-Scope K-V Store。但是对于新手来说，Context 的概念不算非常的直观，这篇文章来带领大家了解一下 Context 包的基本作用和使用方法。

## 1 包的引入
 在 go1.7 及以上版本 context 包被正式列入官方库中，所以我们只需要import "context"就可以了，而在 go1.6 及以下版本，我们要 import "golang.org/x/net/context"

## 2. Context 基本数据结构
Context interface
Context interface 是最基本的接口

type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
 * Deadline()返回一个time.Time，是当前 Context 的应该结束的时间，ok 表示是否有 deadline
 * Done()返回一个struct{}类型的只读 channel
 * Err()返回 Context 被取消时的错误
 * Value(key interface{}) 是 Context 自带的 K-V 存储功能


 canceler interface
 canceler interface 定义了提供 cancel 函数的 context，当然要求数据结构要同时实现 Context interface

 Structs
除了以上两个 interface 之外，context 包中还定义了若干个struct，来实现上面的 interface

emptyCtx
emptyCtx是空的Context，只实现了Context interface，只能作为 root context 使用。

type emptyCtx int
cancelCtx
cancelCtx继承了Context并实现了cancelerinterface，从WithCancel()函数产生

type cancelCtx struct {
    Context

    done chan struct{} // closed by the first cancel call.

    mu       sync.Mutex
    children map[canceler]bool // set to nil by the first cancel call
    err      error             // set to non-nil by the first cancel call
}
timerCtx
timerCtx继承了cancelCtx，所以也自然实现了Context和canceler这两个interface，由WithDeadline()函数产生

type timerCtx struct {
    cancelCtx
    timer *time.Timer // Under cancelCtx.mu.

deadline time.Time
}
valueCtx
valueCtx包含key、val field，可以储存一对键值对，由WithValue()函数产生

type valueCtx struct {
    Context
    key, val interface{}
}
3. Context 实例化和派生
Context 只定义了 interface，真正使用时需要实例化，官方首先定义了一个 emptyCtx struct 来实现 Context interface，然后提供了Backgroud()函数来便利的生成一个 emptyCtx 实例。
实现代码如下

// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
}

func (*emptyCtx) Done() <-chan struct{} {
    return nil
}

func (*emptyCtx) Err() error {
    return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
    return nil
}

func (e *emptyCtx) String() string {
    switch e {
    case background:
        return "context.Background"
    case todo:
        return "context.TODO"
    }
    return "unknown empty Context"
}

var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)

func Background() Context {
    return background
}
Backgroud() 生成的 emptyCtx 实例是不能取消的，因为emptyCtx没有实现canceler interface，要正常取消功能的话，还需要对 emptyCtx 实例进行派生。常见的两种派生用法是WithCancel() 和 WithTimeout。

WithCancel
调用WithCancel()可以将基础的 Context 进行继承，返回一个cancelCtx示例，并返回一个函数，可以在外层直接调用cancelCtx.cancel()来取消 Context
代码如下：

func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    c := newCancelCtx(parent)
    propagateCancel(parent, &c)
    return &c, func() { c.cancel(true, Canceled) }
}
// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
    return cancelCtx{
        Context: parent,
        done:    make(chan struct{}),
    }
}
WithTimeout
调用WithTimeout，需要传一个超时时间。来指定过多长时间后超时结束 Context，源代码中可以得知WithTimeout是WithDeadline的一层皮，WithDeadline传的是具体的结束时间点，这个在工程中并不实用，WithTimeout会根据运行时的时间做转换。
源代码如下：

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))

func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
    if cur, ok := parent.Deadline(); ok && cur.Before(deadline) {
        // The current deadline is already sooner than the new one.
        return WithCancel(parent)
    }
    c := &timerCtx{
        cancelCtx: newCancelCtx(parent),
        deadline:  deadline,
    }
    propagateCancel(parent, c)
    d := deadline.Sub(time.Now())
    if d <= 0 {
        c.cancel(true, DeadlineExceeded) // deadline has already passed
        return c, func() { c.cancel(true, Canceled) }
    }
    c.mu.Lock()
    defer c.mu.Unlock()
    if c.err == nil {
        c.timer = time.AfterFunc(d, func() {
            c.cancel(true, DeadlineExceeded)
        })
    }
    return c, func() { c.cancel(true, Canceled) }
}
在WithDeadline中，将 timeCtx.timer 挂上结束时的回调函数，回调函数的内容是调用cancel来结束 Context。

WithValue
WithValue的具体使用方法在下面的用例中会讲。
源代码如下：

func WithValue(parent Context, key, val interface{}) Context {
    if key == nil {
        panic("nil key")
    }
    if !reflect.TypeOf(key).Comparable() {
        panic("key is not comparable")
    }
    return &valueCtx{parent, key, val}
}
4. 实际用例
（1）超时结束示例
我们起一个本地的 http serice，名字叫"lazy"，这个 http server 会随机的发出一些慢请求，要等6秒以上才返回，我们使用这个程序来模拟我们的被调用方 hang 住的情况

package main

import (
    "net/http"
    "math/rand"
    "fmt"
    "time"
)


func lazyHandler(w http.ResponseWriter, req *http.Request) {
    ranNum := rand.Intn(2)
    if ranNum == 0 {
        time.Sleep(6 * time.Second)
        fmt.Fprintf(w, "slow response, %d\n", ranNum)
        fmt.Printf("slow response, %d\n", ranNum)
        return
    }
    fmt.Fprintf(w, "quick response, %d\n", ranNum)
    fmt.Printf("quick response, %d\n", ranNum)
    return
}

func main() {
    http.HandleFunc("/", lazyHandler)
    http.ListenAndServe(":9200", nil)
}
然后我们写一个主动调用的 http service，他会调用我们刚才写的"lazy"，我们使用 context，来解决超过2秒的慢请求问题，如下代码：

package main

import (
    "context"
    "net/http"
    "fmt"
    "sync"
    "time"
    "io/ioutil"
)

var (
    wg sync.WaitGroup
)

type ResPack struct {
    r *http.Response
    err error
}

func work(ctx context.Context) {
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    defer wg.Done()
    c := make(chan ResPack, 1)

    req, _ := http.NewRequest("GET", "http://localhost:9200", nil)
    go func() {
        resp, err := client.Do(req)
        pack := ResPack{r: resp, err: err}
        c <- pack
    }()

    select {
    case <-ctx.Done():
        tr.CancelRequest(req)
        <-c
        fmt.Println("Timeout!")
    case res:= <-c:
        if res.err != nil {
            fmt.Println(res.err)
            return
        }
        defer res.r.Body.Close()
        out, _ := ioutil.ReadAll(res.r.Body)
        fmt.Printf("Server Response: %s", out)
    }
    return
}


func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
    defer cancel()
    wg.Add(1)
    go work(ctx)
    wg.Wait()
    fmt.Println("Finished")
}
在 main 函数中，我们定义了一个超时时间为2秒的 context，传给真正做事的work()，work接收到这个 ctx 的时候，需要等待 ctx.Done() 返回，因为 channel 关闭的时候，ctx.Done() 会受到空值，当 ctx.Done()返回时，就意味着 context 已经超时结束，要做一些扫尾工作然后 return 即可。

（2）使用 WithValue 制作生成 Request ID 中间件
在 Golang1.7 中，"net/http"原生支持将Context嵌入到 *http.Request中，并且提供了http.Request.Conext() 和 http.Request.WithContext()这两个函数来新建一个 context 和 将 context 加入到一个http.Request实例中。下面的程序演示了一下利用WithValue()创建一个可以储存 K-V 的 context，然后写一个中间件来自动获取 http头部的 "X-Rquest-ID"值，加入到 context 中，使业务函数可以直接取到该值。

package main

import (
    "net/http"
    "context"
    "fmt"
)

const requestIDKey = "rid"

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
    reqID := req.Header.Get("X-Request-ID")
    if reqID == "" {
        reqID = "0"
    }
    return context.WithValue(ctx, requestIDKey, reqID)
}

func requestIDFromContext(ctx context.Context) string {
    return ctx.Value(requestIDKey).(string)
}

func middleWare(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        ctx := newContextWithRequestID(req.Context(), req)
        next.ServeHTTP(w, req.WithContext(ctx))
    })
}

func h(w http.ResponseWriter, req *http.Request) {
    reqID := requestIDFromContext(req.Context())
    fmt.Fprintln(w, "Request ID: ", reqID)
    return
}

func main() {
    http.Handle("/", middleWare(http.HandlerFunc(h)))
    http.ListenAndServe(":9201", nil)
}

https://yq.aliyun.com/articles/69662?spm=5176.8067842.tagmain.29.xiqYFT




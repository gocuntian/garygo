package main

import (
    "fmt"
    "net/http"
    "context"
)
//context
// Variables
// var Canceled = errors.New("context canceled")
// Cancelled是当上下文被取消时Context.Err返回的错误。

// var DeadlineExceeded error = deadlineExceededError{}
// DeadlineExceeded是Context.Err在上下文的最后期限过去时返回的错误。
// type CancelFunc func()
// type Context interface {
//     Deadline() (deadline time.Time, ok bool)
//     Done() <-chan struct{} //返回一个struct{}类型的只读 channel
//     Err() error
//     Value(key interface{}) interface{}
// }
// func Background() Context
// 返回一个非空的上下文。 它从不取消，没有值，没有期限。 它通常由主函数，初始化和测试使用，并作为传入请求的顶级上下文。
// func TODO() Context
// TODO返回一个非空的上下文。 代码应该使用context.TODO当它不清楚使用哪个上下文或它还不可用时（因为周围函数尚未扩展为接受上下文参数）。 TODO被静态分析工具识别，该工具确定上下文是否在程序中正确传播。
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
// WithCancel返回具有新的完成通道的父级的副本。 返回的上下文的完成通道在调用返回的cancel函数时或当父上下文的完成通道关闭时（以先发生者为准）关闭。
// 取消此上下文释放与其相关联的资源，因此代码应在此上下文中运行的操作完成后立即调用cancel。

// func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
// WithDeadline返回父上下文的副本，其中最后期限调整为不晚于d。 如果父级的截止日期早于d，则WithDeadline（parent，d）在语义上等同于父级。 返回的上下文的完成通道在截止时间到期时，调用返回的cancel函数时，或者当父上下文的完成通道关闭时（以先发生者为准）关闭。
// 取消此上下文释放与其相关联的资源，因此代码应在此上下文中运行的操作完成后立即调用cancel。

// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithTimeout返回WithDeadline（parent，time.Now（）。Add（timeout））。
// 取消此上下文释放与其关联的资源，因此代码应在此上下文中运行的操作完成时立即调用cancel：

// func WithValue(parent Context, key, val interface{}) Context




//Go 1.7 添加了 context 包，用于传递数据和做超时、取消等处理。
//*http.Request 添加了 r.Context() 和 r.WithContext() 来操作请求过程需要的 context.Context 对象。
//func (r *Request) Context() context.Context {}
//func (r *Request) WithContext(ctx context.Context) *Request {}
// 传递数据
// context 可以在 http.HandleFunc 之间传递数据：

func handleone(w http.ResponseWriter, r *http.Request){
    ctx:=context.WithValue(r.Context(),"key","abcdef") //写入string 到 context
    handletwo(w,r.WithContext(ctx)) //// 传递给下一个 handleFunc
}

func handletwo(w http.ResponseWriter, r *http.Request){
    str, ok := r.Context().Value("key").(string)//// 取出的 interface 需要推断到 string
    if !ok{
        str = "no string"
    }
    w.Write([]byte("context.key = " + str))
}

func main(){
    http.HandleFunc("/",handleone)
    if err := http.ListenAndServe(":8888",nil); err != nil{
        fmt.Println("start http server fail:",err)
    }
}

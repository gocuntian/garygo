package main

import (
    "log"
    "net/http"
)


/*
|=============================================================================================================================================
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
实现了Handler接口的对象可以注册到HTTP服务端，为特定的路径及其子树提供服务。
ServeHTTP应该将回复的头域和数据写入ResponseWriter接口然后返回。返回标志着该请求已经结束，HTTP服务端可以转移向该连接上的下一个请求。

type HandlerFunc func(ResponseWriter, *Request)
HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。
如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口。

func (HandlerFunc) ServeHTTP
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
ServeHTTP方法会调用f(w, r)

|=============================================================================================================================================
|---------------------------------------------------------------------------------------------------------------------------------------------------
func Handle(pattern string, handler Handler)
Handle注册HTTP处理器handler和对应的模式pattern（注册到DefaultServeMux）。如果该模式已经注册有一个处理器，Handle会panic。ServeMux的文档解释了模式的匹配机制

func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
HandleFunc注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）。ServeMux的文档解释了模式的匹配机制。
|---------------------------------------------------------------------------------------------------------------------------------------------------
*/

func middlewareOne(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        log.Println("Executing middlewareOne")
        next.ServeHTTP(w,r)
        log.Println("Executing middlewareOne again")
    })
}

func middlewareTwo(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        log.Println("Executing middlewareTwo")
        if r.URL.Path != "/" {
            return  //注意
        }
        next.ServeHTTP(w,r)
        log.Println("Executing middlewareTwo again")
    })
}

func final(w http.ResponseWriter, r *http.Request){
    log.Println("Executing finalHander")
    w.Write([]byte("OK"))
}

func main(){
    finalHander := http.HandlerFunc(final)
    http.Handle("/",middlewareOne(middlewareTwo(finalHander)))
    http.ListenAndServe(":3000",nil)
}
/*
http://localhost:3000/
2017/03/27 09:48:43 Executing middlewareOne
2017/03/27 09:48:43 Executing middlewareTwo
2017/03/27 09:48:43 Executing finalHander

2017/03/27 09:48:43 Executing middlewareTwo again
2017/03/27 09:48:43 Executing middlewareOne again
==============================================================
http://localhost:3000/info
2017/03/27 09:48:13 Executing middlewareOne
2017/03/27 09:48:13 Executing middlewareTwo

2017/03/27 09:48:13 Executing middlewareOne again
*/



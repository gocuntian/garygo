package main

import (
    "fmt"
    "net/http"
)

type MyHandler struct{} //实现 http.Handler 接口的 ServeHTTP 方法

func (mh MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
    if r.URL.Path == "/hello"{
        w.Write([]byte("hello"))
        return
    }
    if r.URL.Path == "/xingcuntian"{
        w.Write([]byte("xingcuntian"))
        return
    }
    fmt.Println(r.URL.Path)
    w.Write([]byte("index"))
}

func main(){
    //1
    // http.Handle("/",MyHandler{})
    // if err:=http.ListenAndServe(":12345",nil);err!=nil{
    //     fmt.Println("start http server fail:",err)
    // }

    if err:=http.ListenAndServe(":12345",MyHandler{});err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
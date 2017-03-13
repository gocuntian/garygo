package main

import (
    "fmt"
    "net/http"
    "context"
)

func handle1(w http.ResponseWriter,r *http.Request){
    ctx:=context.WithValue(r.Context(),"name","xingcuntian")// 写入 string 到 context
    handle2(w, r.WithContext(ctx))  // 传递给下一个 handleFunc
}

func handle2(w http.ResponseWriter, r *http.Request){
    str, ok:=r.Context().Value("name").(string)// 取出的 interface 需要推断到 string
    if !ok{
        str = "not string"
    }
    w.Write([]byte("context.name = "+str))
}

func main(){
    http.HandleFunc("/",handle1)
    if err:=http.ListenAndServe(":8088",nil);err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
package main

import (
    "fmt"
    "net/http"
)

func handle1(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("handle1"))
}

func handle2(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("handle2"))
}

func makeHandlers(handlers ...http.HandlerFunc) http.HandlerFunc{
    fmt.Println(handlers)
    return func(w http.ResponseWriter, r *http.Request){
        for _,handler:=range handlers {
            handler(w,r)
        }
    }
}

func main(){
    http.HandleFunc("/", makeHandlers(handle1, handle2))
    if err:=http.ListenAndServe(":12345",nil);err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
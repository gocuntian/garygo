package main

import (
    "fmt"
    "net/http"
)

func main(){
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("hello, Go HTTP Server"))
    })

    http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Hi,Go HTTP hello"))
    })

    if err:=http.ListenAndServe(":12345",nil);err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
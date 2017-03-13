package main

import (
    "fmt"
    "net/http"
)

func HttpHandle(w http.ResponseWriter, r *http.Request){
  fmt.Println("Method:",r.Method)
  fmt.Println("URL: ",r.URL, "URL.Path :",r.URL.Path)
  fmt.Println("RemoteAddress: ",r.RemoteAddr)
  fmt.Println("UserAgent: ",r.UserAgent())
  fmt.Println("Header.Accept: ",r.Header.Get("Accept"))
  fmt.Println("Cookies: ",r.Cookies())
}

func main(){
    http.HandleFunc("/",HttpHandle)
    if err:=http.ListenAndServe(":12345",nil);err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
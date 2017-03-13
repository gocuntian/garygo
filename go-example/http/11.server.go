package main

import (
    "net/http"
    "path"
)

func main(){
    http.HandleFunc("/",handle)
    http.ListenAndServe(":3000",nil)
}

func handle(w http.ResponseWriter,r *http.Request){
    fp:=path.Join("images","foo.png")
    http.ServeFile(w,r,fp)
}
package main

import (
    "net/http"
)

func main(){
    http.HandleFunc("/",handle)
    http.ListenAndServe(":3000",nil)
}

func handle(w http.ResponseWriter,r *http.Request){
    w.Header().Set("Server","A Go Web Server")
    w.WriteHeader(200)
}
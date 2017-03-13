package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"hello HTTP Server")
}

func main(){
    mux:=http.NewServeMux()
    mux.HandleFunc("/",handler)
    http.ListenAndServe(":12345",mux)
}
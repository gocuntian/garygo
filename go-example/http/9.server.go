package main

import "net/http"

func main(){
    http.HandleFunc("/",Handle)
    http.ListenAndServe(":3000",nil)
}

func Handle(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("OK"))
}
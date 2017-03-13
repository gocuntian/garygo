package main

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func handle(w http.ResponseWriter,r *http.Request,_ httprouter.Params){
    
    fmt.Fprint(w,"hello, httprouter")
}

func main(){
    router:=httprouter.New()
    router.GET("/",handle)

    if err:=http.ListenAndServe(":12345",router);err!=nil{
        fmt.Println("start http server fail:",err)
    }
}
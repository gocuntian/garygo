package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/kataras/iris.v6/middleware/pprof"
)

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.Get("/",func(ctx *iris.Context){
        ctx.HTML(iris.StatusOK,"<h1> Please click <a href='/debug/pprof'>here</a>")
    })
    app.Get("/debug/pprof/*action",pprof.New())
    app.Listen(":9080")
}
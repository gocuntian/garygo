package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.Favicon("./static/favicons/iris_favicon_32_32.ico")
    app.Get("/",func(ctx *iris.Context){
        ctx.HTML(iris.StatusOK,`this is test`)
    })
    app.Listen(":8088")
}
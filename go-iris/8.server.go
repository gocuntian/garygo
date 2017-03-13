package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.StaticWeb("/static","./assets")
    app.Listen(":8088")
}
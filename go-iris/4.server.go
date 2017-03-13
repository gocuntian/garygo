package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/kataras/iris.v6/adaptors/view"
)

type Todo struct{
    Task string
    Done bool
}

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.Adapt(view.HTML("./templates",".html"))
    todos:=[]Todo{
        {"Learn Go",true},
        {"Read GopherBook",true},
        {"Create a web app in Go",false},
    }

    app.Get("/",func(ctx *iris.Context){
        ctx.Render("todos.html",struct{Todos []Todo}{todos})
    })
    app.Listen(":8088")
}
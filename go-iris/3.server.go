package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/gorillamux"
)

func main(){
    app:=iris.New()
    app.Adapt(gorillamux.New())/// Adapt the "httprouter", you can use "gorillamux" too.
    userAges:=map[string]int{
        "Alice":23,
        "Bob":26,
        "Claire":34,
    }
    app.Get("/users/{name}",func(ctx *iris.Context){
        name:=ctx.Param("name")
        age:=userAges[name]
        ctx.Writef("%s is %d years old!",name,age)
    })
    app.Listen(":8088")
}
package main

import (
    "context"
    "time"
    "gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main(){
    app:=iris.New()
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.Get("/hi",func(ctx *iris.Context){
        ctx.HTML(iris.StatusOK," <h1>hi, I just exist in order to see if the server is closed</h1>")
    })
    app.Adapt(iris.EventPolicy{
        Interrupted:func(*iris.Framework){
           ctx, _ := context.WithTimeout(context.Background(),5 * time.Second)
           app.Shutdown(ctx)
        },
    })

    app.Listen(":8099")
}
package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/kataras/iris.v6/adaptors/view"
)

type ContactDetails struct{
    Email string
    Subject string
    Message string
}

func main(){
    app:=iris.New(iris.Configuration{Gzip:false,Charset:"UTF-8"})
    app.Adapt(iris.DevLogger())
    app.Adapt(httprouter.New())
    app.Adapt(view.HTML("./templates",".html"))

    app.Get("/",func(ctx *iris.Context){
        ctx.Render("forms.html",nil)
    })

    app.Post("/",func(ctx *iris.Context){
        // details:=ContactDetails{
        //     Email:ctx.FormValue("email"),
        //     Subject:ctx.FormValue("subject"),
        //     Message:ctx.FormValue("message"),
        // }
        
        var details ContactDetails
        ctx.ReadForm(&details)
        _=details
        ctx.Render("forms.html",struct{Success bool}{true})
    })

    app.Listen(":8088")
}
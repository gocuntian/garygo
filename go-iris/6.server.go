package main

import (
    //"fmt"
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/kataras/iris.v6/adaptors/view"
)

type ContactDetails struct{
    Email string `json:"email"`
    Subject string `json:"subject"`
    Message string  `json:"message"`
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
        //     Email:ctx.FormValue("Email"),
        //     Subject:ctx.FormValue("Subject"),
        //     Message:ctx.FormValue("Message"),
        // }
        
        var details ContactDetails
        err := ctx.ReadForm(&details)
		if err != nil {
			ctx.Log(iris.DevMode, "Error when reading form: "+err.Error())
		}


        // details := ContactDetails{}
		// err := ctx.ReadForm(&details)
		// if err != nil {
		// 	ctx.Log(iris.DevMode, "Error when reading form: "+err.Error())
		// }
        ctx.JSON(iris.StatusOK,details)
       // ctx.Writef("details: %#v", details)
       // ctx.Render("forms.html",struct{Success bool}{true})
    })

    app.Listen(":8088")
}
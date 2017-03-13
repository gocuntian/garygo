package main

import (
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/httprouter"
    "gopkg.in/kataras/iris.v6/adaptors/sessions"
)

var (
    key = "my_sessionid"
)

func secret(ctx *iris.Context){
    if auth,_:=ctx.Session().GetBoolean("authenticated");!auth{
        ctx.EmitError(iris.StatusForbidden)
        return
    }
    ctx.WriteString("The cake is lie!")
}

func login(ctx *iris.Context){
    session:=ctx.Session()
    session.Set("authenticated",true)
}

func logout(ctx *iris.Context){
    session:=ctx.Session()
    session.Set("authenticated",false)
}

func main(){
    app:=iris.New()
    app.Adapt(httprouter.New())
    sess:=sessions.New(sessions.Config{Cookie:key})
    app.Adapt(sess)
    app.Get("/secret",secret)
    app.Get("/login",login)
    app.Get("logout",logout)
    app.Listen(":8088")
}
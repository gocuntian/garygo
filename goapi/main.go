package main

import (
    "gopkg.in/macaron.v1"
    "github.com/xingcuntian/go_test/goapi/routers"
    "github.com/xingcuntian/go_test/goapi/modules/middleware"
	"github.com/xingcuntian/go_test/goapi/routers/account"
)

func main(){
    routers.GlobalInit()
    m:=macaron.Classic()
    m.Use(macaron.Renderer())
    m.Use(middleware.Contexter())
    m.Get("/",myhandler)
    m.Group("/account",func(){
        m.Post("/create",account.Create)
        m.Get("/info",account.Info)
    })
    m.Run()
}

func myhandler(ctx *middleware.Context){
    ctx.ErrorJSON(401,"error message")
    return
}
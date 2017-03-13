package main

import (
   // "fmt"
    "gopkg.in/macaron.v1"
    "github.com/Unknwon/paginater"
)

func main(){
    m:=macaron.Classic()
    m.Use(macaron.Renderer())

    m.Get("/",func(ctx *macaron.Context){
        ctx.Data["Name"] = "xingcuntian"
        page := ctx.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pager := paginater.New(100,10, page, 5)

    
	ctx.Data["Page"] = pager
       // fmt.Println(p)
        ctx.HTML(200,"index")
    })

    m.Run()
}
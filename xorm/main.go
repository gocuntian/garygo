package main

import (
   // "fmt"
    "strconv"
    "gopkg.in/macaron.v1"
)

func main(){
    m:=macaron.Classic()
    m.Get("/add/:username",AddHandler)
   // m.Get("/list",ListHandler)
    m.Get("/update/:id/:name",UpdateHandler)
    m.Get("/delete/:id",DeleteHandler)
    m.Run()
}

func AddHandler(ctx *macaron.Context) string{
    name:= ctx.Params(":username")
    if len(name) > 0{
        err:=CreateAccount(name)
        if err!=nil{
           return "create user fail" 
        }
        return "create user success"
    }else{
        return "create user fail"
    }
}

// func ListHandler()string{
//    _,err:=ListAccount()
//    if err!=nil{
//        return "select data fail"
//    }
//    return "ok"
// //    for i,one:=range list{
// //       return fmt.Sprintf("%d %#v\n",i+1,one)
// //    }
// }

func UpdateHandler(ctx *macaron.Context) string{
    id:=ctx.Params(":id")
    Numid,err:=strconv.ParseInt(id,10,64)
    if err!=nil{
        return "id change fail"
    }
    name:=ctx.Params(":name")
    err=UpdateAccount(name,Numid)
    if err!=nil{
        return "update fail"
    }
    return "update success"+name
}

func DeleteHandler(ctx *macaron.Context)string{
    id:=ctx.Params(":id")
    Numid,err:=strconv.ParseInt(id,10,64)
    if err!=nil{
        return "id change fail"
    }
    err=DeleteAccount(Numid)
    if err!=nil{
        return "delete data fail"
    }
    return "delete data success"
}
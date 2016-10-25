package main

import (
    "gopkg.in/macaron.v1"
)

func main(){
    m:=macaron.Classic()
    m.Get("/add/:username",AddHandler)
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
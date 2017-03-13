package main

import "github.com/go-martini/martini"


type Person struct{
    name string
}

func (p Person) say_hi() string{
    return "hello"+p.name
}

func main(){
    m:=martini.Classic()
    p:=Person{name:"xingcuntian"}
    m.Get("/",p.say_hi)
    m.Run()
}
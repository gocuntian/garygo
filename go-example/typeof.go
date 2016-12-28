package main

import (
    "fmt"
)
//Golang里面获取类型可以用reflect或者fmt
//1)fmt.Printf("%T\n",str)
//2)fmt.Println(reflect.TypeOf(str))
//使用
// str:=fmt.Sprintf("%T\n",str)//获取的是字符串表示的类型
// str:=reflect.TypeOf(str) //获得的是类型
//还可以定义 interface，然后通过switch配合(type)来判断
func main(){
    var a interface{}
    a ="1"
    switch vtype:=a.(type){
        case string:
            fmt.Println("string")
        case int:
            fmt.Println("int")
        default:
            fmt.Println(vtype)
    }
}
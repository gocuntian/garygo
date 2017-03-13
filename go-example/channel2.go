package main

import "fmt"


func main(){
    ci:= make(chan int,1)
    go write(ci,4)
    go write(ci,5)
    go write(ci,6)

    value:= <-ci
    value1:= <-ci
    value2:= <-ci

    fmt.Println(value,value1,value2)
}

func write(c chan int,num int){
  c<-num
}

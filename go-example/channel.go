package main

import "fmt"


func main(){
    ci:= make(chan int)
    ci <- 4
    // ci <- 5
    // ci <- 6

    // value:= <-ci
    // value1:= <-ci
    value2:= <-ci
 fmt.Println(value2)
    // fmt.Println(value,value1,value2)
}

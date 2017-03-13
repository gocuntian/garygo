package main

import "fmt"

func main(){
    queue := make(chan string,2)
    queue <- "one this is one string"
    queue <- "two this is two string"
    close(queue)
    for elem := range queue{
        fmt.Println(elem)
    }
}
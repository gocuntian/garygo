package main

import (
    "fmt"
    "context"
)

type favContextKey string

func main(){
    f := func(ctx context.Context, k favContextKey){
        if v := ctx.Value(k); v != nil {
            fmt.Println("found value: ", v)
            return
        }
        fmt.Println("key not found:", k)
    }
    k := favContextKey("language")

    ctx := context.WithValue(context.Background(),k,"Golang")

    f(ctx,k)
    f(ctx,favContextKey("lang"))

}



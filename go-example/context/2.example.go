package main

import (
    "fmt"
    "time"
    "context"
)
//Example(WithDeadline)


func main(){
    d := time.Now().Add(50 * time.Millisecond)
   // d := time.Now().Add(2 * time.Second)
    ctx, cancel := context.WithDeadline(context.Background(),d)
    defer cancel()

    select{
        case <-time.After(1 * time.Second):
            fmt.Println(<-ctx.Done())
            fmt.Println("overslept")
        case <-ctx.Done():
            fmt.Println(ctx)
            fmt.Println(<-ctx.Done())
            fmt.Println(ctx.Err())
    }
}
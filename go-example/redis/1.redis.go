package main

import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
)
/*
如果我们留意Redis的回复的话，我们发现Resp对象有很多有用的函数让Go的类型转换变得容易：
Resp.Bytes() – converts a single reply to a byte slice ([]byte)
Resp.Float64() – converts a single reply to a Float64
Resp.Int() – converts a single reply to a int
Resp.Int64() – converts a single reply to a int64
Resp.Str() – converts a single reply to a string
Resp.Array() – converts an array reply to an slice of individual Resp objects ([]*Resp)
Resp.List() – converts an array reply to an slice of strings ([]string)
Resp.ListBytes() – converts an array reply to an slice of byte slices ([][]byte)
Resp.Map() – converts an array reply to a map of strings, using each item in the array reply alternately as the keys and values for the map (map[string]string)
*/
func main(){
    conn,err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    //授权
    resp := conn.Cmd("auth","sensetime@2016")
    
    fmt.Println(resp)
    if resp.Err != nil {
        log.Fatal(resp.Err)
    }

    // resp = conn.Cmd("HMSET","album:4","title","Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
    // fmt.Println(resp.Str())
    // if resp.Err != nil {
    //     log.Fatal(resp.Err)
    // }

    err = conn.Cmd("HMSET","album:5","title","Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8).Err
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Electric Ladyland added!")
}
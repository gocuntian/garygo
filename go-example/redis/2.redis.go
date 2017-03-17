package main

import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
    "log"
)
// 如果我们留意Redis的回复的话，我们发现Resp对象有很多有用的函数让Go的类型转换变得容易：

// Resp.Bytes() – converts a single reply to a byte slice ([]byte)
// Resp.Float64() – converts a single reply to a Float64
// Resp.Int() – converts a single reply to a int
// Resp.Int64() – converts a single reply to a int64
// Resp.Str() – converts a single reply to a string
// Resp.Array() – converts an array reply to an slice of individual Resp objects ([]*Resp)
// Resp.List() – converts an array reply to an slice of strings ([]string)
// Resp.ListBytes() – converts an array reply to an slice of byte slices ([][]byte)
// Resp.Map() – converts an array reply to a map of strings, using each item in the array reply alternately as the keys and values for the map (map[string]string)

func main(){
    conn, err := redis.Dial("tcp","localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    
    defer conn.Close()

    //授权
    err = conn.Cmd("auth","sensetime@2016").Err
    if err !=nil {
        log.Fatal(err)
    }

    title, err := conn.Cmd("HGET","album:1","title").Str()
    if err != nil {
        log.Fatal(err)
    }

    artist, err := conn.Cmd("HGET","album:1","artist").Str()
    if err != nil {
        log.Fatal(err)
    }

    price, err := conn.Cmd("HGET","album:1","price").Float64()
    if err != nil {
        log.Fatal(err)
    }

    likes, err := conn.Cmd("HGET","album:1","likes").Int()
    if err != nil {
        log.Fatal(err)
    } 

    fmt.Printf("%s by %s: £%.2f [%d likes]\n", title, artist, price, likes)

    info, err := conn.Cmd("HGETALL","album:1").List()
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(info)

    info2, err := conn.Cmd("HGETALL","album:1").Map()
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(info2)

    info3, err := conn.Cmd("HGETALL","album:1").ListBytes()
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(info3)

    info4, err := conn.Cmd("HGETALL","album:1").Array()
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(info4)

}
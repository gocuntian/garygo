package main

import (
	"fmt"
	"gopkg.in/redis.v5"
)

var client *redis.Client

func newClient(){
    client=redis.NewClient(&redis.Options{
      Addr: "127.0.0.1:6379",
      Password: "sensetime@2016",
      DB:   0,
    })
    pong,err:=client.Ping().Result()
    fmt.Println(pong,err)
}

func testClient(){
   err:=client.Set("name","xingcuntian",0).Err() 
   if err!=nil{
     panic(err)
   }

   val,err:=client.Get("name").Result()
   if err!=nil{
      panic(err)
   }
   fmt.Println("name :",val)

   val2,err:=client.Get("age").Result()
   if err == redis.Nil{
      fmt.Println("age does not exists")
   }else if err!=nil{
      panic(err)
   }else{
      fmt.Println("age",val2)
   }
}

func main(){
    newClient()
    testClient()
}

package main

import (
	"fmt"
	"github.com/xingcuntian/go_test/go-example/rpc/client"
	"github.com/xingcuntian/go_test/go-example/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
	reply2 := client.GetUserByIdRequest(c)
	fmt.Println(reply2.Name)
}
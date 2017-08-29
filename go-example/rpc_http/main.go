package main

import (
	"fmt"
	"os"

	"github.com/xingcuntian/go_test/go-example/rpc_http/client"
	"github.com/xingcuntian/go_test/go-example/rpc_http/server"
)

func main() {
	if len(os.Args) > 1 {
		c := client.CreateClient()
		defer c.Close()
		reply := client.PerformRequest(c)
		fmt.Println(reply.Message)
	} else {
		server.StartServer()
	}

}

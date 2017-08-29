package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/xingcuntian/go_test/go-example/rpc_http/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(c *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse
	if err := c.Call("HelloWorldHandler.HelloWorld", args, &reply); err != nil {
		log.Fatal("error:", err)
	}

	return reply
}

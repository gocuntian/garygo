package main

import (
	"fmt"

	"github.com/xingcuntian/go_test/go-example/rpc_http_json/client"
	"github.com/xingcuntian/go_test/go-example/rpc_http_json/contract"
)

func main() {
	fmt.Println("start")
	var response contract.HelloWorldResponse
	response = client.PerformRequest()
	fmt.Println(response)

}

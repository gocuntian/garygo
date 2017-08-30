package main

import "github.com/xingcuntian/go_test/go-example/rpc_http_json/server"

// curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234
func main() {
	server.StartServer()
}

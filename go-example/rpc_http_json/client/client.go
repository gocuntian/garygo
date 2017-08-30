package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc/json"
	"github.com/xingcuntian/go_test/go-example/rpc_http_json/contract"
)

func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
	defer r.Body.Close()
	fmt.Println(r.Body)
	var result contract.HelloWorldResponse
	if err := json.DecodeClientResponse(r.Body, &result); err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	return result
}

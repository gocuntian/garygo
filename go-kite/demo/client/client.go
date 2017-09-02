package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("second", "1.0.0")

	client := k.NewClient("http://localhost:6000/kite")
	client.Dial()

	response, _ := client.Tell("kite.ping")
	fmt.Println(response.MustString())
}

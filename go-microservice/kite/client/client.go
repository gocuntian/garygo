package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("exp22", "1.0.0")
	client := k.NewClient("http://localhost:6000/kite")
	client.Dial()

	response, _ := client.Tell("Hello", "Nic")
	fmt.Println(response.MustString())
}

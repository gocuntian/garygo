package main

import (
	"fmt"

	"github.com/koding/kite"
	"github.com/koding/kite/protocol"
)

func main() {
	k := kite.New("second", "1.0.0")
	kites, _ := k.GetKites(&protocol.KontrolQuery{
		Username:    k.Config.Username,
		Environment: k.Config.Environment,
		Name:        "first",
	})
	fmt.Println(kites)
	// client := kites[0]

	// client.Dial()
	// response, _ := client.Tell("hello", 9)
	// fmt.Println(response.MustFloat64())
}

package main

import (
	"fmt"
	"os"

	"github.com/json-iterator/go"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteVal(group)
	if stream.Error != nil {
		fmt.Println("error:", stream.Error)
		return
	}
	os.Stdout.Write(stream.Buffer())
	fmt.Println()
}

package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","PHP","Maroon"]}`)
	fmt.Printf(jsoniter.Get(val, "Colors", 2).ToString())
	fmt.Println()
}

package main

import (
	"fmt"
	"log"

	"github.com/json-iterator/go"
)

type Animal struct {
	Name  string
	Order string
}

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)
	var animals []Animal
	err := jsoniter.Unmarshal(jsonBlob, &animals)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v", animals)
}

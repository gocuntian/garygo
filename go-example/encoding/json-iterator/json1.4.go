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
	iter := jsoniter.ConfigFastest.BorrowIterator(jsonBlob)
	defer jsoniter.ConfigFastest.ReturnIterator(iter)
	iter.ReadVal(&animals)
	if iter.Error != nil {
		log.Println(iter.Error)
		return
	}
	fmt.Printf("%+v", animals)
	fmt.Println()
}

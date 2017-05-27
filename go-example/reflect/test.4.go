package main

import (
	"fmt"
)

type Song struct {
	Name   string `json:"name" xml:"name"`
	Length int    `json:"length" xml:"length"`
}

func main() {
	var songInterface interface{} = Song{Name: "测试", Length: 126}
	if _, ok := songInterface.(Song); ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

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
	switch songInterface.(type) {
	case Song:
		fmt.Println("Yes, I am song type !")
	default:
		fmt.Println("No, I am not song type!")
	}
}

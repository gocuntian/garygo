package main

import (
	"fmt"
	"log"
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
	b, err := jsoniter.Marshal(group)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(b)

	os.Stdout.Write(b)
	fmt.Println()
}

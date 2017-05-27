package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	greetSymbol, err := p.Lookup("Greet")
	if err != nil {
		panic(err)
	}

	greet := greetSymbol.(func() string)
	fmt.Println(greet())
}

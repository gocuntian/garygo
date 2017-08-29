package main

import "fmt"

//go:generate echo hello
//go:generate go run generate.go
//go:generate echo file=$GOFILE pkg=$GOPACKAGE

func main() {
	fmt.Println("main func")
}

// go generate
// hello
// main func
// file=generate.go pkg=main

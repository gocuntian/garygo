package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog.......")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat.....")
}

func main() {
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Println("start server err")
	}
}

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
	//func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }
	http.Handle("/dog", d)
	http.Handle("/cat", c)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println("start server err")
	}
}

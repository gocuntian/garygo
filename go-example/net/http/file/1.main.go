package main

import (
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}

func main() {
	http.HandleFunc("/foo", foo)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}

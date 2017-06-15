package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("error ListenAndServe")
	}
}

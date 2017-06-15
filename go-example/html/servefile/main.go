package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is test")
}

func bar(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

func chiem(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "small.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/chiem", chiem)
	log.Println("start server:9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}

}

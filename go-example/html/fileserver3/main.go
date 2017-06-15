package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func dogs(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main() {
	//fs := http.FileServer(http.Dir("public"))
	//http.Handle("/pics/", fs)
	http.Handle("/resources/",http.StripPrefix("/resources",http.FileServer(http.Dir("public"))))
	http.HandleFunc("/dogs", dogs)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}
//http://localhost:9090/resources/pics/dog.jpeg

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Println(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"
	if err := tpl.ExecuteTemplate(w, "index.html", person{f, l, s}); err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}

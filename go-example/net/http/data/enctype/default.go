package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	Firstname  string
	Lastname   string
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
	//body
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	fmt.Println(bs)
	body := string(bs)
	fmt.Println(body)
	err := tpl.ExecuteTemplate(w, "index.html", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

//BODYS: first=ddd&last=gffgggg&subscribe=on

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/bowzer", bowzer)
	http.HandleFunc("/dog/bowzer/pictures", bowzerpics)
	http.HandleFunc("/cat", cat)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	var c *http.Cookie
	c, err := r.Cookie("uname")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func bowzer(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:  "uname",
		Value: "xingcuntian",
		Path:  "/",
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "bowzer.gohtml", c)
}

func bowzerpics(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("uname")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(w, "bowzerpics.gohtml", c)
}

func cat(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("uname")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	tpl.ExecuteTemplate(w, "cat.gohtml", c)
}

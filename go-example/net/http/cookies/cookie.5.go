package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users/goofy", set)
	http.HandleFunc("/users/goofy/query", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username3")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Println(c)
	fmt.Fprintln(w, c)
}

func set(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:  "username3",
		Value: "xingcuntian3",
	}
	http.SetCookie(w, c)
	fmt.Println(c)
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username3")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "you cookie : ", c)
}

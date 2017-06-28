package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	if err := http.ListenAndServe(":9090",nil); err != nil {
		fmt.Println(err)
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{
	    Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400),http.StatusBadRequest)
		return 
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
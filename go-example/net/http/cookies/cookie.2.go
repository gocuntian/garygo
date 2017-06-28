package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/abundance",abundance)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	if err := http.ListenAndServe(":9090",nil); err != nil {
		log.Println(err)
	}
}


func set(w http.ResponseWriter, r *http.Request ) {
	http.SetCookie(w,&http.Cookie{
		Name:"username",
		Value:"xingcuntian",
	})
	fmt.Fprintln(w,"Cookie written")
}

func read(w http.ResponseWriter, r *http.Request ) {
	c1, err := r.Cookie("username")
	if err != nil {
		log.Println(err)
	}else{
		fmt.Fprintln(w, "Your Cookie #1:",c1)
	}

	c2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	}else{
		fmt.Fprintln(w, "Your Cookie #2:",c2)
	}

	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", c3)
	}
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:"general",
		Value:"this is two",
	})
	http.SetCookie(w,&http.Cookie{
		Name:"specific",
		Value:"this is three",
	})
	fmt.Fprintln(w,"this is ok")
}


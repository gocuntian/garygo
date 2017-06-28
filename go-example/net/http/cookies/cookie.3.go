package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username2")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "username2",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

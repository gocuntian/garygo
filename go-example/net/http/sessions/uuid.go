package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

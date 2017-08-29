package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go web app powered by Docker")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	if err := http.ListenAndServe(":18080", nil); err != nil {
		log.Println(err)
	}
}

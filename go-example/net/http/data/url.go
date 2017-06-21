package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	fmt.Fprintln(w, "Do search: "+v)
}

//http://localhost:9090/?q=9999

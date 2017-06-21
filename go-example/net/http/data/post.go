package main

import (
	"io"
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	  <input type="text" name = "q">
	  <input type="submit">
	</form>
	<br/>`+v)
}

package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func cat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dog(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	defer f.Close()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/toby.jpg", dog)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/",foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	if err := http.ListenAndServe(":9090",nil); err != nil {
		log.Println(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
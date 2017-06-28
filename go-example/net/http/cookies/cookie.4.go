package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.HandleFunc("/expire",expire)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	if err := http.ListenAndServe(":9090",nil); err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,`<h1><a href="/set">to set cookie</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{
		Name:"username2",
		Value: "xingcuntian2",
	})
	fmt.Fprintln(w,"set cookie is success!")
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username2")
	if err != nil {
		http.Redirect(w,r,"/set",http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username2")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w,c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
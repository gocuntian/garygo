package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func loginHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w,r)
        log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
    })
}

func index(w http.ResponseWriter, r *http.Request){
    log.Println("Execute index Handler")
    fmt.Fprintf(w,"About, Go lang!")
}

func about(w http.ResponseWriter, r *http.Request){
    log.Println("Execute about Handler")
    fmt.Fprintf(w,"Welcome!")
}

func icoHandler(w http.ResponseWriter, r *http.Request){
   
}

func main(){
    http.HandleFunc("/favicon.icon",icoHandler)
    indexHandler := http.HandlerFunc(index)
    aboutHandler := http.HandlerFunc(about)

    http.Handle("/",loginHandler(indexHandler))
    http.Handle("/about",loginHandler(aboutHandler))

    server := &http.Server{
        Addr: ":3200",
    }
    log.Println("listening...")
    server.ListenAndServe()
}
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.Handle("/", loggingHandler(http.HandlerFunc(index)))
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Print("Index Handler started")
	defer log.Print("Indx handler ended")
	log.Print(r.Context().Value("uname"))
	ctx := r.Context()
	select {
	case <-time.After(4 * time.Second):
		fmt.Fprintln(w, "hello Gopher")
	case <-ctx.Done():
		err := ctx.Err()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		ctx := context.WithValue(r.Context(), "uname", "xingcuntian")
		next.ServeHTTP(w, r.WithContext(ctx))
		log.Printf("Complated %s in %v", r.URL.Path, time.Since(start))
	})
}

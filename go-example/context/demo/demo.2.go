package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "user_id", 888)
	ctx = context.WithValue(ctx, "user_name", "xingcuntian")
	results := dbAccess(ctx)
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("user_id").(int)
	return uid
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	fmt.Println("=============================")
	fmt.Fprintln(w, ctx)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	//"errors"
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
	results, err := dbAccess(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		//context deadline exceeded
		return
	}
	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ch := make(chan int)
	go func() {
		uid := ctx.Value("user_id").(int)
		time.Sleep(10 * time.Second)
		if ctx.Err() != nil {
			return
		}
		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
		//return 0, errors.New("ddddddd")
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx)
	fmt.Println("=============================")
	fmt.Fprintln(w, ctx)
}

package main

import (
	"net/http"
	"runtime/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}

// func Lookup(name string) *Profile
// Lookup返回具有指定名字的Profile；如果没有，会返回nil。
//http://localhost:11181/
// goroutine profile: total 10007
// 10000 @ 0x42d0ea 0x42d1ce 0x4065d1 0x406255 0x643472 0x457831
// #	0x643471	main.f+0x41	/data/go/src/github.com/xingcuntian/go_test/go-example/runtime/pprof/pprof.2.go:11

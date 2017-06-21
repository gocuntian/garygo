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
	ctx := r.Context()
	log.Println(ctx)
	fmt.Println("=============================")
	fmt.Fprintln(w, ctx)
}

// 2017/06/20 15:25:02 context.Background.WithValue(
// 	&http.contextKey{name:"http-server"},
// 	&http.Server{
// 		Addr:":9090", Handler:http.Handler(nil), TLSConfig:(*tls.Config)(0xc4200e4000),
// 		ReadTimeout:0, ReadHeaderTimeout:0, WriteTimeout:0, IdleTimeout:0, MaxHeaderBytes:0,
// 		TLSNextProto:map[string]func(*http.Server, *tls.Conn, http.Handler){"h2":(func(*http.Server, *tls.Conn, http.Handler))(0x5f7440)},
// 		ConnState:(func(net.Conn, http.ConnState))(nil), ErrorLog:(*log.Logger)(nil), disableKeepAlives:0, inShutdown:0,
// 		nextProtoOnce:sync.Once{m:sync.Mutex{state:0, sema:0x0},done:0x1},
// 		nextProtoErr:error(nil), mu:sync.Mutex{state:0, sema:0x0},
// 		listeners:map[net.Listener]struct {}{http.tcpKeepAliveListener{TCPListener:(*net.TCPListener)(0xc42000e038)}:struct {}{}},
// 		activeConn:map[*http.conn]struct {}{(*http.conn)(0xc420096780):struct {}{}}, doneChan:(chan struct {})(nil)
//    }

// ).WithValue(
// 	  &http.contextKey{name:"local-addr"},
// 	  &net.TCPAddr{IP:net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, Port:9090, Zone:""}
// 	  ).WithCancel.WithCancel

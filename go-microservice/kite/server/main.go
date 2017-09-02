package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"

	"github.com/koding/kite"
)

func main() {
	k := kite.New("math", "1.0.0")
	//c := config.MustGet()
	//k.Config = c
	k.Config.Port = 6000
	//k.Config.KontrolURL = "http://kontrol:6000/kite"
	k.Config.DisableAuthentication = true
	//k.RegisterForever(&url.URL{Scheme: "http", Host: "127.0.0.1:8091", Path: "/kite"})
	k.HandleFunc("Hello", func(r *kite.Request) (interface{}, error) {
		name, err := r.Args.One().String()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		return fmt.Sprintf("Hello %v", name), nil
	})
	k.Config.Port = 8091
	k.Run()
}

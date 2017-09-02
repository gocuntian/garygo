package main

import "github.com/koding/kite"

func main() {
	k := kite.New("first", "1.0.0")
	k.Config.Port = 6000
	k.Config.DisableAuthentication = true
	k.HandleFunc("hello", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		return a * a, nil
	})
	k.Run()
}

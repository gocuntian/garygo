package main

import (
	"sync"
	"fmt"
)

func main() {
	list := map[string]interface{}{
		"name": "xingcuntian",
		"birthday": "1990.10.14",
		"age":27,
		"hobby": []string{"xingcuntian","age","ddddsdsdsdsd"},
		"constellation":"gary",
	}

	var m sync.Map
	for k, v := range list{
		m.Store(k,v)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		m.Store("age",22)
		m.LoadOrStore("tag",888)
		wg.Done()
	}()

	go func() {
		m.Delete("constellation")
		m.Store("age",18)
		wg.Done()
	}()

	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})

}

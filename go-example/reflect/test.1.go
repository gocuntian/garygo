package main

import (
	"fmt"
	"reflect"
)

type Blog struct {
	Name string
}

func (this Blog) Test() string {
	fmt.Println("this is Test method")
	return this.Name
}

func main() {
	var o interface{} = &Blog{"xingcuntian"}
	v := reflect.ValueOf(o)
	fmt.Println(v)
	m := v.MethodByName("Test")
	rets := m.Call([]reflect.Value{})
	fmt.Println(rets)
	fmt.Println(rets[0])
}
// &{xingcuntian}
// this is Test method
// [xingcuntian]
// xingcuntian


package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("lib.so")
	if err != nil {
		panic(err)
	}

	add, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	fmt.Println(add.(func(a, b int) int)(1, 2))
}

// type Plugin Uses

// type Plugin struct {
//     // contains filtered or unexported fields
// }
// 插件是一个加载的Go插件。

// func Open Uses
// func Open(path string) (*Plugin, error)
// 打开打开一个Go插件。 如果路径已经打开，则返回现有的*插件。 多个goroutine并发使用是安全的。

// func (*Plugin) Lookup Uses
// func (p *Plugin) Lookup(symName string) (Symbol, error)
// Lookup在插件p中搜索名为symName的符号。 符号是任何导出的变量或函数。 如果找不到符号，则会报告错误。 多个goroutine并发使用是安全的。

// type Symbol Uses
// type Symbol interface{}
// 符号是指向变量或函数的指针。

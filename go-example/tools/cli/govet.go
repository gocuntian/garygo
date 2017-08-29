package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://golang.org")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, res.Body)
}

//go vet govet.go

// 命令go vet是一个用于检查Go语言源码中静态错误的简单工具。与大多数Go命令一样，go vet命令可以接受-n标记和-x标记。-n标记用于只打印流程中执行的命令而不真正执行它们。
// -n标记也用于打印流程中执行的命令，但不会取消这些命令的执行。

// func main() {
// 	res, err := http.Get("https://golang.org")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	io.Copy(os.Stdout, res.Body)
// }

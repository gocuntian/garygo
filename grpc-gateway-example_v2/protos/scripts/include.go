package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fs, _ := ioutil.ReadDir(".")
	out, _ := os.Create("swagger.pb.go")
	out.Write([]byte("package protos \n\nconst(\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			name := strings.TrimPrefix(f.Name(), "service.")
			out.Write([]byte(strings.TrimSuffix(name, ".json") + "= `"))
			f, _ := os.Open(f.Name())
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}

package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	addr      = flag.String("addr", ":8080", "http server address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("github.com/xingcuntian/go_test/go-example/net/websocket/demo2", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

// func Import(path, srcDir string, mode ImportMode) (*Package, error)
// // 导入是Default.Import的缩写。
// build.FindOnly
// 如果设置了FindOnly，则在找到应包含包的源的目录之后，导入将停止。 它不会读取目录中的任何文件。
func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTempl.Execute(w, r.Host)
}

func main() {
	flag.Parse()
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "home.html")))
	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}

// func Must(t *Template, err error) *Template
// Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化：
// func ParseFiles(filenames ...string) (*Template, error)
// ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义。返回的模板的名字是第一个文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。至少要提供一个文件。如果发生错误，会停止解析并返回nil。
// func Join(elem ...string) string
// Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符。结果是经过简化的，所有的空字符串元素会被忽略。

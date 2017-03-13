package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

// 获取大小的借口
type Size interface {
    Size() int64
}

type Stat interface {
    Stat() (os.FileInfo, error)
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
        r.ParseMultipartForm(32 << 20) // 32m
        file, _, err := r.FormFile("userfile")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        defer file.Close()
        f, err := os.Create("filenametosaveas")
        defer f.Close()
        io.Copy(f, file)
        // 获取文件信息的接口
        if statInterface, ok := file.(Stat); ok {
            fileInfo, _ := statInterface.Stat()
            fmt.Fprintf(w, "a上传文件的大小为: %d", fileInfo.Size())
        }
        if sizeInterface, ok := file.(Size); ok {
            fmt.Fprintf(w, "b上传文件的大小为: %d", sizeInterface.Size())
        }
        return
    }
    // 上传页面
    w.Header().Add("Content-Type", "text/html")
    w.WriteHeader(200)
    html := `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="userfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`
    io.WriteString(w, html)
}
func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
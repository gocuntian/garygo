package main

import (
    "fmt"
    "net/http"
)

func HttpHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello1"))
    // hj, ok := w.(http.Hijacker)
    // if !ok {
    //     return
    // }
    // conn, buf, err := hj.Hijack()
    // if err != nil {
    //     w.WriteHeader(500)
    //     return
    // }
    // defer conn.Close()       // 需要手动关闭连接
    // w.Write([]byte("hello")) // 会提示 http: response.Write on hijacked connection

    // // 返回内容需要
    // buf.WriteString("hello")
    // buf.Flush()
}

func main(){
    http.HandleFunc("/hj",HttpHandle)
    if err := http.ListenAndServe(":50000",nil); err != nil {
        fmt.Println("start http server fail:", err)
    }
}
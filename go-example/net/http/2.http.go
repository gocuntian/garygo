package main

import (
    "log"
    "bytes"
    "net/http"
)

// func StatusText(code int) string
// StatusText返回HTTP状态码code对应的文本，如220对应"OK"。如果code是未知的状态码，会返回""。

// func Error(w ResponseWriter, error string, code int)
// Error使用指定的错误信息和状态码回复请求，将数据写入w。错误信息必须是明文。

// func DetectContentType(data []byte) string
// DetectContentType函数实现了http://mimesniff.spec.whatwg.org/描述的算法，用于确定数据的Content-Type。函数总是返回一个合法的MIME类型；
// 如果它不能确定数据的类型，将返回"application/octet-stream"。它最多检查数据的前512字节。

func enforceXMLHandler(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if r.ContentLength == 0 {
            http.Error(w, http.StatusText(400),400)
            return
        }
        buf := new(bytes.Buffer)
        buf.ReadFrom(r.Body)
        if http.DetectContentType(buf.Bytes()) != "text/xml; charset=utf-8" {
            http.Error(w, http.StatusText(415),415)
            return
        }
        next.ServeHTTP(w,r)
    })
}

func final(w http.ResponseWriter, r *http.Request) {
  log.Println("Executing finalHandler")
  w.Write([]byte("OK"))
}

func main(){
   finalHander := http.HandlerFunc(final)
   http.Handle("/",enforceXMLHandler(finalHander))
   http.ListenAndServe(":3000",nil)
}
package main

import (
    "os"
    "log"
    "compress/gzip"
)
//压缩文件
func main(){
    outputFile, err := os.Create("test.txt.gz")
    if err != nil {
        log.Fatal(err)
    }
    
    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    _, err = gzipWriter.Write([]byte("Gohpers\n"))
    if err != nil {
        log.Fatal(err)
    }
    log.Println("compressed data written to file.")
}

// func NewWriter(w io.Writer) *Writer
// NewWriter创建并返回一个Writer。写入返回值的数据都会在压缩后写入w。
//调用者有责任在结束写入后调用返回值的Close方法。因为写入的数据可能保存在缓冲中没有刷新入下层。

// 如要设定Writer.Header字段，调用者必须在第一次调用Write方法或者Close方法之前设置。
//Header字段的Comment和Name字段是go的utf-8字符串，但下层格式要求为NUL中止的ISO 8859-1 (Latin-1)序列。
//如果这两个字段的字符串包含NUL或非Latin-1字符，将导致Write方法返回错误。
package main

import (
    "os"
    "io"
    "log"
    "compress/gzip"
)
//文件解压
func main(){
    gzipFile, err := os.Open("test.txt.gz")
    if err != nil {
        log.Fatal(err)
    }
    defer gzipFile.Close()

    gzipReader, err := gzip.NewReader(gzipFile)
    if err != nil {
        log.Fatal(err)
    }

    outfileWriter, err := os.Create("unzipped.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer outfileWriter.Close()

    _, err = io.Copy(outfileWriter,gzipReader)
    if err != nil {
        log.Fatal(err)
    }
}
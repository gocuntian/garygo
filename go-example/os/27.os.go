package main

import (
    "os"
    "io"
    "log"
    "net/http"
)
//http文件下载
func main(){
    newFile, err := os.Create("devdungeon.html")
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()

    url := "http://www.csdn.net/article/2012-07-05/2807113-less-is-exponentially-more"
    response, err := http.Get(url)
    defer response.Body.Close()

    numBytesWritten, err := io.Copy(newFile,response.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Downloaded %d byte file.\n",numBytesWritten)
}
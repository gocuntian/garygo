package main

import (
    "os"
    "log"
)
//写入字节流到文件
func main(){
    file, err := os.OpenFile("test.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE,0666)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    byteSlice :=[]byte("Bytes!\n")
    bytesWritten, err := file.Write(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("wrote %d bytes.\n",bytesWritten)
}
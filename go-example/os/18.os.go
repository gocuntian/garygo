package main

import (
    "os"
    "log"
    "io"
)
//最少读取多少个字节
func main(){
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    byteSlice := make([]byte,512)
    minBytes := 8
    numBytesRead, err := io.ReadAtLeast(file,byteSlice,minBytes)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n",numBytesRead)
    log.Printf("Data read: %s\n",byteSlice)
}
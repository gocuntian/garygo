package main

import (
    "log"
    "os"
)

func main(){
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    byteSlice :=make([]byte,16)

    bytesReadNum, err := file.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Number of bytes read: %d\n",bytesReadNum)
    log.Printf("Data read: %s\n",byteSlice)
}
package main

import (
    "fmt"
    "log"
    "io"
    "os"
    "crypto/md5"
)

func main(){
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    hasher := md5.New()
    _, err = io.Copy(hasher,file)
    if err != nil {
        log.Fatal(err)
    }
    sum := hasher.Sum(nil)
    fmt.Printf("Md5 checksum: %x\n",sum)
}
//Md5 checksum: 66613d903d08fd4ec0e1f0b75bba6720
package main

import (
    "log"
    "os"
)
//Truncate a File(文件截取)
func main(){
    err := os.Truncate("newtest.txt",100)
    if err != nil {
        log.Fatal(err)
    }
}
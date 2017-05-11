package main

import (
    "log"
    "os"
)

//Delete a File(文件删除)
func main(){
    err := os.Remove("new")
    if err != nil {
        log.Fatal(err)
    }
}
package main

import (
    "log"
    "os"
)
//Rename and Move a File(移动和重命名文件)

func main(){
    originalPath := "newtest.txt"
    newPath := "text.txt"
    err :=os.Rename(originalPath,newPath)
    if err != nil{
        log.Fatal(err)
    }
}
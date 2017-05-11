package main

import (
    "image"
    "image/png"
    "os"
    "log"
)
//将图像写入文件
func main(){
    myImage := image.NewRGBA(image.Rect(0,0,100,200))

    outputFile, err := os.Create("test.png")
    if err != nil {
        log.Fatal(err)
    }
    defer outputFile.Close()

    png.Encode(outputFile,myImage)
}
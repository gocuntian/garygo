package main

import (
    "encoding/base64"
    "image"
    "image/png"
    "fmt"
    "bytes"
)
//Base64编码的图像
func main(){
    myImage := image.NewRGBA(image.Rect(0,0,10,20))

    var buf bytes.Buffer

    png.Encode(&buf,myImage)

    encodeString := base64.StdEncoding.EncodeToString(buf.Bytes())

    htmlImage := "<img src=\"data:image/png;base64," +encodeString + "\">"

    fmt.Println(htmlImage)
}
package main

import (
    "fmt"
    "log"
    "encoding/base64"
    "os"
    "io/ioutil"
)

func main(){
    file, err := os.Open("gopher.png")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    buf, err := ioutil.ReadAll(file)
    encodeString := base64.StdEncoding.EncodeToString(buf)
    htmlImage := "<img src=\"data:image/png;base64," +encodeString + "\">"

    fmt.Println(htmlImage)
}

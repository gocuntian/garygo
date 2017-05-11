package main

import (
    "os"
    "log"
    "fmt"
    "io/ioutil"
)
//读取文件所有字节
func main(){
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }

    data, err :=ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Data as hex:%x\n",data)
    fmt.Printf("Data as string: %s\n",data)
    fmt.Println("Number of bytes read:",len(data))
    // Data as hex:7373736173737338
    // Data as string: sssasss8
    // Number of bytes read: 8

}
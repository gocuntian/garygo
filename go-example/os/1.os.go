package main

import (
    "log"
    "os"
)
// type File struct {
//     // 内含隐藏或非导出字段
// }
var (
    newFile *os.File
    err error
)
//Create Empty File(创建空文件)：
func main(){
    newFile, err = os.Create("newtest.txt")
    if err != nil {
        log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
}
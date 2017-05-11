package main

import (
    "fmt"
    "log"
    "os"
)
//Get File Info(获取文件信息)
var (
    fileInfo *os.FileInfo
    err      error
)

func main(){
    fileInfo, err := os.Stat("new")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File name:",fileInfo.Name())
    fmt.Println("Size in bytes:",fileInfo.Size())
    fmt.Println("Permissions:",fileInfo.Mode())
    fmt.Println("Last modified:",fileInfo.ModTime())
    fmt.Println("Is Directory:",fileInfo.IsDir())
    fmt.Printf("System interface type: %T\n",fileInfo.Sys())
    fmt.Printf("System info : %+v\n\n",fileInfo.Sys())
}
// func (p *ProcessState) Sys() interface{}
// Sys返回该已退出进程系统特定的退出信息。需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.WaitStatus类型以获取其内容。
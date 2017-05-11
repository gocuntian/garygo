package main

import (
    "log"
    "io/ioutil"
)
//快速读取文件到内存
// func ReadFile(filename string) ([]byte, error)
// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
func main(){
    data, err := ioutil.ReadFile("test.txt")
    if err !=nil {
        log.Fatal(err)
    }
    log.Printf("Data read : %s\n",data)
}
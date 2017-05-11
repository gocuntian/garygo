package main

import (
    "io/ioutil"
    "log"
)
//快速文件写入
func main(){
    err := ioutil.WriteFile("test.txt",[]byte("Hi\n"),0666)
    if err != nil {
        log.Fatal(err)
    }
}

// func ReadFile(filename string) ([]byte, error)
// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。

// func WriteFile(filename string, data []byte, perm os.FileMode) error
// 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。

// func ReadDir(dirname string) ([]os.FileInfo, error)
// 返回dirname指定的目录的目录信息的有序列表。
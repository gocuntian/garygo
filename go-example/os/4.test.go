package main

import (
    "fmt"
    "io/ioutil"
    "os"
)
// func ReadFile(filename string) ([]byte, error)
// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。

// func WriteFile(filename string, data []byte, perm os.FileMode) error
// 函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
 //将整个文件的内容读到一个字符串里：
//  可以使用 io/ioutil 包里的 ioutil.ReadFile() 方法，该方法第一个返回值的类型是 []byte，里面存放读取到的内容，
//  第二个返回值是错误，如果没有错误发生，第二个返回值为 nil。类似的，函数WriteFile() 可以将 []byte 的值写入文件。

func main(){
    dir,_:=os.Getwd()
    outputFile:=dir+"/readme_new.md"

    if is_exists,_:=os.Stat(outputFile);  is_exists == nil{
        _, _ = os.Create(outputFile)
    }

    inputFile := dir+"/readme.md"

    buf, err := ioutil.ReadFile(inputFile)
    if err != nil{
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        // panic(err.Error())
    }
    fmt.Printf("%s\n",string(buf))

    err = ioutil.WriteFile(outputFile,buf,0x644)
    if err != nil {
        panic(err.Error())
    }

}
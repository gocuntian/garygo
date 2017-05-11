package main

import (
    "log"
    "os"
    "archive/zip" //zip包提供了zip档案文件的读写服务
)
//文件归档

func main(){
    outFile, err := os.Create("test.zip")
    if err != nil{
        log.Fatal(err)
    }
    defer outFile.Close()

   // func NewWriter(w io.Writer) *Writer
   // NewWriter创建并返回一个将zip文件写入w的*Writer。

   zipWriter :=zip.NewWriter(outFile)

    var filesToArchive = []struct{ 
            Name string
            Body string
         }{
             {"test.txt","this is test1"},
             {"text.txt","this is test2"},
         }
        
    for _, file :=range filesToArchive {
        
        fileWriter, err :=zipWriter.Create(file.Name)
        if err !=nil {
            log.Fatal(err)
        }
        _,err = fileWriter.Write([]byte(file.Body))
        if err != nil {
            log.Fatal(err)
        }
    }
    err = zipWriter.Close()
    if err != nil {
        log.Fatal(err)
    }
    

}

// func (w *Writer) Create(name string) (io.Writer, error)
// 使用给出的文件名添加一个文件进zip文件。本方法返回一个io.Writer接口（用于写入新添加文件的内容）。
// 文件名必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔。新增文件的内容必须在下一次调用CreateHeader、
// Create或Close方法之前全部写入
// func (w *Writer) Close() error
// Close方法通过写入中央目录关闭该*Writer。本方法不会也没办法关闭下层的io.Writer接口。
package main

import (
    "os"
    "log"
    "io"
)
//文件复制
func main(){
    originalFile, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer originalFile.Close()

    newFile, err := os.Create("test_copy.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()

    bytesWritten, err := io.Copy(newFile,originalFile)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Copied %d bytes.",bytesWritten)

    err = newFile.Sync()
    if err != nil {
        log.Fatal(err)
    }
}

// func Copy(dst Writer, src Reader) (written int64, err error)
// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

// 对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。

// func (f *File) Sync() (err error)
// Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存。
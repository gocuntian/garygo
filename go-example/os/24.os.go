package main

import (
    "log"
    "os"
    "io"
    "path/filepath"
    "archive/zip"
)
//解压缩归档文件
func main(){
    zipReader, err := zip.OpenReader("test.zip")
    if err != nil {
        log.Fatal(err)
    }
    defer zipReader.Close()
    
    for _, file :=range zipReader.Reader.File{
        zippedFile, err := file.Open()
        if err != nil {
            log.Fatal(err)
        }
        defer zippedFile.Close()

        targetDir := "./test"
        extractedFilePath :=filepath.Join(
            targetDir,
            file.Name,
        )

        if file.FileInfo().IsDir() {
            log.Println("creating directory:", extractedFilePath)
            os.MkdirAll(extractedFilePath,file.Mode())
        }else{
            log.Println("extracting file:",file.Name)
            outputFile, err := os.OpenFile(
                extractedFilePath,
                os.O_WRONLY|os.O_CREATE|os.O_TRUNC,file.Mode(),
            )
            if err != nil {
                log.Fatal(err)
            }
            defer outputFile.Close()
            _, err =io.Copy(outputFile,zippedFile)
            if err != nil {
                log.Fatal(err)
            }
        }
    }

}

// func OpenReader(name string) (*ReadCloser, error)
// OpenReader会打开name指定的zip文件并返回一个*ReadCloser。

// type Reader struct {
//     File    []*File
//     Comment string
//     // 内含隐藏或非导出字段
// }

// func (h *FileHeader) FileInfo() os.FileInfo
// FileInfo返回一个根据h的信息生成的os.FileInfo。

// type FileInfo interface {
//     Name() string       // 文件的名字（不含扩展名）
//     Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
//     Mode() FileMode     // 文件的模式位
//     ModTime() time.Time // 文件的修改时间
//     IsDir() bool        // 等价于Mode().IsDir()
//     Sys() interface{}   // 底层数据来源（可以返回nil）
// }

// func (m FileMode) IsDir() bool
// IsDir报告m是否是一个目录。



// func Join(elem ...string) string
// Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加路径分隔符。结果是经过简化的，所有的空字符串元素会被忽略。

// func Copy(dst Writer, src Reader) (written int64, err error)
// 将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

// 对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。
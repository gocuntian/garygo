package main

import (
    "log"
    "os"
)
//Check if File Exists(检查文件是否存在)
var (
    fileInfo *os.FileInfo
    err      error
)


func main(){
    fileInfo, err := os.Stat("text.txt")
    if err != nil {
        if os.IsNotExist(err){
            log.Fatal("File does not exist.")
        }
    }
    log.Println("File does exist. File information:")
    log.Println(fileInfo)
}

// type FileInfo interface {
//     Name() string       // 文件的名字（不含扩展名）
//     Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
//     Mode() FileMode     // 文件的模式位
//     ModTime() time.Time // 文件的修改时间
//     IsDir() bool        // 等价于Mode().IsDir()
//     Sys() interface{}   // 底层数据来源（可以返回nil）
// }
// FileInfo用来描述一个文件对象。

// func Stat(name string) (fi FileInfo, err error)
// Stat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。如果出错，返回的错误值为*PathError类型。

// func Lstat(name string) (fi FileInfo, err error)
// Lstat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接。如果出错，返回的错误值为*PathError类型。

// func IsNotExist(err error) bool
// 返回一个布尔值说明该错误是否表示一个文件或目录不存在。ErrNotExist和一些系统调用错误会使它返回真。
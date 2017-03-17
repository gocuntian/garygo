package main

import (
    "fmt"
    "os"
    "time"
)
// go 标准库 os.File
/*
#2. File 结构体

### 1)func Create(name string) (file *File, err error) // Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
### 2)func Open(name string) (file *File, err error) // Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式
### 3)func (f *File) Stat() (fi FileInfo, err error) // Stat返回描述文件f的FileInfo类型值
### 4)func (f *File) Readdir(n int) (fi []FileInfo, err error) // Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序
### 5)func (f *File) Read(b []byte) (n int, err error) // Read方法从f中读取最多len(b)字节数据并写入b
### 6)func (f *File) WriteString(s string) (ret int, err error) // 向文件中写入字符串
### 7)func (f *File) Sync() (err error) // Sync递交文件的当前内容进行稳定的存储。一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存
### 8)func (f *File) Close() error // Close关闭文件f，使文件不能用于读写

###func Stat(name string) (fi FileInfo, err error) // Stat 返回描述文件的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接
###func Lstat(name string) (fi FileInfo, err error) // Lstat 返回描述文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接。
*/
func main(){
    //获取当前目录
    dir, err := os.Getwd()
    fmt.Println(dir,err)

    file :=dir+"/new"
    var fh *os.File
    fi, _ :=os.Stat(file)
    if fi == nil{
        fh, _ = os.Create(file)//文件不存在就创建
    }else{
        fh, _ = os.OpenFile(file,os.O_RDWR,0666)//打开文件
    }

    ret, err :=fh.WriteString("ddddddd")
    fmt.Println(ret,err)

    w := []byte("hello go langage" + time.Now().String())
    n, err := fh.Write(w)
    fmt.Println(n,err)

    // 设置下次读写位置
    rets, err := fh.Seek(0,0)
    fmt.Println("当前文件指针位置",rets,err)

    b := make([]byte,128)
    n, err = fh.Read(b)
    fmt.Println(n,err,string(b))
    fh.Close()

    var dh os.FileInfo
    newDir := dir+"/dirs"
    dh, _ = os.Stat(newDir)

    if dh == nil{
        err := os.MkdirAll(newDir,0755)
        if err != nil{
            fmt.Println(err)
        }
        dh, _ = os.Stat(newDir)
    } 
    fmt.Println(dh.Name())
    fmt.Println(dh.Size())
    fmt.Println(dh.Mode())
    fmt.Println(dh.ModTime())
    fmt.Println(dh.IsDir())
    fmt.Println(dh.Sys())
}
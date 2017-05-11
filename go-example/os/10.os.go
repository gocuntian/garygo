package main

import (
    "os"
    "log"
    "fmt"
)
//硬连接和符号连接
// func Symlink(oldname, newname string) error
// Symlink创建一个名为newname指向oldname的符号链接。如果出错，会返回* LinkError底层类型的错误。

// func Link(oldname, newname string) error
// Link创建一个名为newname指向oldname的硬链接。如果出错，会返回* LinkError底层类型的错误。

// func Lstat(name string) (fi FileInfo, err error)
// Lstat返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接。如果出错，返回的错误值为*PathError类型。
// func Lchown(name string, uid, gid int) error
// Chmod修改name指定的文件对象的用户id和组id。如果name指定的文件是一个符号链接，它会修改该符号链接自身的用户id和组id。如果出错，会返回*PathError底层类型的错误。
func main(){
    err := os.Link("original_also.txt","original_l.txt")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("creating sym")
    err = os.Symlink("original_also.txt","original_al_syl.txt")
    if err != nil {
        log.Fatal(err)
    }

    fileInfo, err := os.Lstat("original_al_syl.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("link info: %+v",fileInfo)

    err = os.Lchown("original_al_syl.txt",os.Getuid(),os.Getgid())
    if err != nil {
        log.Fatal(err)
    }

}
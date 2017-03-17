package main

import (
    "fmt"
    "os"
)
/*
#1. os 常用导出函数

### 1)func Hostname() (name string, err error) // Hostname返回内核提供的主机名
### 2)func Environ() []string // Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝
### 3)func Getenv(key string) string // Getenv检索并返回名为key的环境变量的值
### 4)func Getpid() int // Getpid返回调用者所在进程的进程ID
### 5)func Exit(code int) // Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行
### 6)func Stat(name string) (fi FileInfo, err error) // 获取文件信息
### 7)func Getwd() (dir string, err error) // Getwd返回一个对应当前工作目录的根路径
### 8)func Mkdir(name string, perm FileMode) error // 使用指定的权限和名称创建一个目录
### 9)func MkdirAll(path string, perm FileMode) error // 使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误
### 10)func Remove(name string) error // 删除name指定的文件或目录
### 11)func TempDir() string // 返回一个用于保管临时文件的默认目录
### 12)var Args []string Args保管了命令行参数，第一个是程序名。
*/
func main(){
    //预定义变量，保存命令行参数
    fmt.Println(os.Args)
    // 获取host name
    fmt.Println(os.Hostname())
    fmt.Println(os.Getpid())

    // 获取全部环境变量
    env := os.Environ()
    for k, v := range env {
        fmt.Println(k,v)
    }

       // 终止程序
    // os.Exit(1)
    
    // 获取一条环境变量
    fmt.Println(os.Getenv("PATH"))
    fmt.Println(os.Getenv("GOPATH"))

     // 获取当前目录
     dir, err := os.Getwd()
     fmt.Println(dir,err)

     // 创建目录一级
    //  err = os.Mkdir(dir+"/new",0755)
    //  fmt.Println(err)

      
    //  fi, err := os.Stat(dir+"/new2")
    //  fmt.Println(fi)
    //  fmt.Println(err)

     // 创建多级目录
    //  err = os.MkdirAll(dir+"/new2/file",0755)
    //  fmt.Println(err)

    // 删除目录
    // err = os.Remove(dir+"/new")
    // fmt.Println(err)
    // err = os.Remove(dir+"/new2")
    // fmt.Println(err)

    //创建临时文件 /tmp
    // tmp_dir := os.TempDir()
    // fmt.Println(tmp_dir);

}
package main

import (
    "fmt"
    "path"
)

// Variables
// func IsAbs(path string) bool
// func Split(path string) (dir, file string)
// func Join(elem ...string) string
// func Dir(path string) string
//Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。


// func Base(path string) string
// func Ext(path string) string
// func Clean(path string) string
// func Match(pattern, name string) (matched bool, err error)

func main(){
    pt := "/www/d/path"
    //判断是否是一个绝对路径
    is_abs :=path.IsAbs(pt)
    fmt.Println(is_abs)

    //将一个文件的路径分割为路径和文件名
   // pf := "/data/golang/src/github.com/xingcuntian/go_test/go-example/map.go"
    pf := "../map.go"
    dir, file :=path.Split(pf)
    fmt.Println(dir,file)

    //将多个字符串合并为一个路径
    dir_join :=path.Join("user","local","bin")
    fmt.Println(dir_join)

    //Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。

    path_dir := path.Dir("/data/www/ServiceAPI")
    fmt.Println(path_dir)
    //Base返回路径的最后一个元素
    path_base := path.Base("/data/www/xingcuntian")
    fmt.Println(path_base)

    //返回path文件扩展名
    path_ext :=path.Ext("../map.go")
    fmt.Println(path_ext)
    //Clean单纯的语法操作返回和path代表同一地址的最短路径
    paths :=[]string{
        "a/c",
        "a//c",
        "a/c/.",
        "a/c/b/..",
        "/../a/c",
        "/../a/b/../././/c",
    }
    for _, p :=range paths {
    fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
    }
    //正则匹配路径
    is_match,err :=path.Match("*.go","/data/golang/src/github.com/xingcuntian/go_test/go-example/map.go")
    fmt.Println(is_match,err)
}
package main

import (
    "fmt"
    "os"
)

func main(){
    f,_:=os.Open("/tmp/xct.txt")
    //从头开始，文件指针偏移100
    f.Seek(100,0)
    buffer:=make([]byte,1024)
    //Read后文件指针也会偏移
    _,err:=f.Read(buffer)
    if err!=nil{
        fmt.Println(err)
        return
    }
    //获取文件指针当前位置
    cur_offset,_:=f.Seek(0,os.SEEK_CUR)
    fmt.Printf("current offset is %d\n",cur_offset)
}
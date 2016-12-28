package main

import (
    "bytes"
    "fmt"
)

func main(){
    //字符串拼接
    var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
    for{
        if piece,ok:=getNextString();ok{
            //Write将s的内容写入缓冲中，如必要会增加缓冲容量。
            //返回n为len(p),err总是nil
            //如果缓冲变得太大，Write会采用错误值ErrTooLarge引发panic
            buffer.WriteString(piece)
        }else{
            break
        }
    }
    fmt.Println("拼接后的结果为==>",buffer.String)
}
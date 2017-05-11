package main

import (
    "fmt"
    "bufio"
    "os"
)
// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，
// 且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
// type Reader struct {
//     // 内含隐藏或非导出字段
// }
//func NewReader(rd io.Reader) *Reader
var inputReader *bufio.Reader
var input string
var err error

func main(){
    inputReader = bufio.NewReader(os.Stdin)
    //inputReader 是一个指向 bufio.Reader 的指针。
    //inputReader := bufio.NewReader(os.Stdin) 这行代码，将会创建一个读取器，并将其与标准输入绑定。
    fmt.Println("Please enter some input: ")
    input, err = inputReader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("The input was: %s\n",input)
}
/**
 ＊ 屏幕是标准输出 os.Stdout；os.Stderr 用于显示错误信息，大多数情况下等同于 os.Stdout
    该函数的实参可以是满足 io.Reader 接口的任意对象（任意包含有适当的 Read() 方法的对象，请参考章节11.8），函数返回一个新的带缓冲的 io.Reader 对象，它将从指定读取器（例如 os.Stdin）读取内容。

    返回的读取器对象提供一个方法 ReadString(delim byte)，该方法从输入中读取内容，直到碰到 delim 指定的字符，然后将读取到的内容连同 delim 字符一起放到缓冲区。

    ReadString 返回读取到的字符串，如果碰到错误则返回 nil。如果它一直读到文件结束，则返回读取到的字符串和io.EOF。如果读取过程中没有碰到 delim 字符，将返回错误 err != nil。

    在上面的例子中，我们会读取键盘输入，直到回车键（\n）被按下。
＊/

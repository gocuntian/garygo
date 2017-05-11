package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)

// func NewReader(rd io.Reader) *Reader
// NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
// func NewWriter(w io.Writer) *Writer
// NewWriter创建一个具有默认大小缓冲、写入w的*Writer。


// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
// ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner。

// ReadLine尝试返回一行数据，不包括行尾标志的字节。如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分。该行剩下的部分将在之后的调用中返回。返回值isPrefix会在返回该行最后一个片段时才设为false。返回切片是缓冲的子切片，只在下一次读取操作之前有效。ReadLine要么返回一个非nil的line，要么返回一个非nil的err，两个返回值至少一个非nil。

// // 返回的文本不包含行尾的标志字节（"\r\n"或"\n"）。如果输入流结束时没有行尾标志字节，方法不会出错，也不会指出这一情况。在调用ReadLine之后调用UnreadByte会总是吐出最后一个读取的字节（很可能是该行的行尾标志字节），即使该字节不是ReadLine返回值的一部分。
// func (b *Writer) WriteString(s string) (int, error)
// WriteString写入一个字符串。返回写入的字节数。如果返回值nn < len(s)，还会返回一个错误说明原因。
func main(){
    dir,_:=os.Getwd()
    inputFile, _ := os.Open(dir+"/readme.md")
    outputFile, _ := os.OpenFile(dir+"/test.txt",os.O_WRONLY|os.O_CREATE,0666)
    defer inputFile.Close()
    defer outputFile.Close()
    inputReader := bufio.NewReader(inputFile)
    outputWriter :=bufio.NewWriter(outputFile)
    for {
        inputString, _, readerError := inputReader.ReadLine()
        fmt.Println(inputString)
        if readerError == io.EOF{
            fmt.Println("EOF")
            return
        }
        outputString :=string([]byte(inputString)[2:5]) + "\r\n"
        fmt.Println(outputString)
        n, err := outputWriter.WriteString(outputString)
        if err != nil{
            fmt.Println(err)
            return
        }
        _ =outputWriter.Flush()
        fmt.Printf("line : %d\n",n)
    }

    fmt.Println("Conversion done")
}

// 一个输入文件 readme.md，然后以每一行为单位读取，从读取的当前行中截取第 3 到第 5 的字节写入另一个文件。然而当你运行这个程序，输出的文件却是个空文件。找出程序逻辑中的 bug，修正它并测试

// func (b *Writer) Flush() error
// Flush方法将缓冲中的数据写入下层的io.Writer接口。

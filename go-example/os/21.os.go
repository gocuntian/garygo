package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
)

func main(){
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    bufferedReader := bufio.NewReader(file)

    byteSlice := make([]byte,5)
    // Peek返回输入流的下n个字节，而不会移动读取位置。返回的[]byte只在下一次调用读取操作前合法。如果Peek返回的切片长度比n小，它也会返会一个错误说明原因。如果n比缓冲尺寸还大，返回的错误将是ErrBufferFull。
    byteSlice, err = bufferedReader.Peek(5)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Peeked at 5 bytes:%s\n",byteSlice)

    numBytesRead, err := bufferedReader.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read %d bytes: %s\n",numBytesRead,byteSlice)

    myByte, err :=bufferedReader.ReadByte()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read 1 byte: %c\n",myByte)

    dataBytes, err := bufferedReader.ReadBytes('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read bytes: %s\n",dataBytes)

    dataString, err := bufferedReader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Read string : %s\n",dataString)

}
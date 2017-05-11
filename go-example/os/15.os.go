package main

import (
    "log"
    "os"
    "bufio"
)

func main(){
    file, err := os.OpenFile("test.txt",os.O_WRONLY,0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    //创建写入缓冲器
    bufferedWriter := bufio.NewWriter(file)

    bytesAvailable := bufferedWriter.Available()
    log.Printf("Available buffer:%d\n",bytesAvailable)

    bytesWritten, err := bufferedWriter.Write(
        []byte{65,66,67},
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Bytes written: %d\n",bytesWritten)

    bytesWritten, err = bufferedWriter.WriteString(
        "Buffered string\n",
    )
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Bytes written: %d\n",bytesWritten)

    unflushedBufferSize := bufferedWriter.Buffered()
    log.Printf("Bytes buffered : %d\n",unflushedBufferSize)

    bytesAvailable = bufferedWriter.Available()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Available buffered:%d\n",bytesAvailable)
    //数据从bufffer保存磁盘
    bufferedWriter.Flush()
    

    bufferedWriter.Reset(bufferedWriter)

    bytesAvailable = bufferedWriter.Available()
    log.Printf("Available buffer:%d\n",bytesAvailable)

    bufferedWriter = bufio.NewWriterSize(
        bufferedWriter,
        8000,
    )

    bytesAvailable = bufferedWriter.Available()
    log.Printf("Available buffer:%d\n",bytesAvailable)
}
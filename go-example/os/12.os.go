package main

import (
    "fmt"
    "os"
    "log"
)
//搜索文件位置

func main(){
    file, _ := os.Open("test.txt")
    defer file.Close()

    var offset int64 = 5

    var whence int = 0
    newPosition, err := file.Seek(offset,whence)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved to 5:",newPosition)

    newPosition, err = file.Seek(-2,1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved back two:",newPosition)

    currentPosition, err := file.Seek(0,1)
    fmt.Println("Current position:",currentPosition)

    newPosition, err = file.Seek(0,0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Position after seeking 0,0:",newPosition)

}
// func (f *File) Seek(offset int64, whence int) (ret int64, err error)
// Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
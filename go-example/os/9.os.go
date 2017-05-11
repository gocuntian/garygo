package main

import (
    "log"
    "os"
    "time"
)
//更改权限，拥有者和时间戳
func main(){
    err := os.Chmod("test.txt",0777)
    if err != nil {
        log.Println(err)
    }

    err = os.Chown("test.txt",os.Getuid(),os.Getgid())
    if err != nil {
        log.Println(err)
    }

    twoDaysFromNow := time.Now().Add(48 * time.Hour)
    lastAccessTime := twoDaysFromNow
    lastModifyTime := twoDaysFromNow
    err = os.Chtimes("test.txt",lastAccessTime,lastModifyTime)
    if err != nil {
        log.Println(err)
    }
}
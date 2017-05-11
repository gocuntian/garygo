package main

import (
    "fmt"
    "io/ioutil"
    "net"
)

func main(){
    conn, err := net.Dial("tcp","time-track.cn:80")
    checkErr(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkErr(err)

    result, err := ioutil.ReadAll(conn)
    checkErr(err)

    fmt.Println(string(result))
}

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}
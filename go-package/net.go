package main

import (
    "fmt"
    "net"
    "bufio"
)

func main(){
    conn,err:=net.Dial("tcp","localhost:80")
    if err!=nil{
        fmt.Printf("dial err")
    }
    fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
    status, err := bufio.NewReader(conn).ReadString('\n')
    fmt.Println(status,err);
}
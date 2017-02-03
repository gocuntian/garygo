package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

func main(){
    const jsonStream = `{"Name":"Ed","Text":"Knock knock"}`

    type Message struct{
        Name,Text string
    }

    var m Message

    //用json.NewDecoder
    dec:=json.NewDecoder(strings.NewReader(jsonStream))
    dec.Decode(&m)
    fmt.Println(m)

    fmt.Printf("%s: %s\n",m.Name,m.Text)


    //用json.Unmarshal
    json.Unmarshal([]byte(jsonStream), &m)
    
    fmt.Println(m)
    fmt.Println("done")
}
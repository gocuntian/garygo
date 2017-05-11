package main

import (
    "fmt"
    "runtime"
    "encoding/xml"
)

type WechatNotifyInfo struct {
    XMLName        xml.Name `xml:"xml"`
    Appid           CDATA  `xml:"appid"`
}

var Text =
    `<xml>
    <appid><![CDATA[wx0f0df4fda4ff1937]]></appid>
    </xml>
    `
type CDATA struct {
    Text string `xml:",cdata"`
}

func main() {
    fmt.Println("version", runtime.Version())

    msg := &WechatNotifyInfo{}
    err := xml.Unmarshal([]byte(Text), msg)
    if err != nil{
        fmt.Println(err)
    }

    fmt.Println(fmt.Sprintf("%#v", msg))
}
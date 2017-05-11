package main

import (
    "encoding/xml"
    "fmt"
    "runtime"
)

type WechatNotifyInfo struct{
    XMLName xml.Name `xml:"xml"`
    Appid   string `xml:"appid,CDATA"`//cdata改成大写  go1.2.1小写可以
}

var Text = 
    `<xml>
       <appid><![CDATA[wx0f0df4fda4ff1937]]></appid>
    </xml>
    `
func main(){
    fmt.Println("version",runtime.Version())
    msg:=&WechatNotifyInfo{}
    err := xml.Unmarshal([]byte(Text),msg)
    if err!=nil {
        fmt.Println(err)
    }
    fmt.Println(fmt.Sprintf("%#v",msg))
}
    
package main

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "workspace@sensetime.com")
	m.SetHeader("To", "xingcuntian@sensetime.com")
	m.SetAddressHeader("Cc", "xingcuntian@qq.com", "xingcuntian")
	m.SetHeader("Subject", "Hello world")
	m.SetBody("text/html", "Hello <b>Gary</b> and <i>xingcuntian</i>!")
	m.Attach("cube.png")
	d := gomail.NewDialer("smtp.partner.outlook.cn", 587, "workspace@sensetime.com", "xxxx")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		fmt.Println("send message is ok")
	}
}

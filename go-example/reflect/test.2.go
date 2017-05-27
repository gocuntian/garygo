package main

import (
	"reflect"
	"fmt"
)

type Song struct {
	Name 	string `json:"name" xml:"name"`
	Length 	int    `json:"length" xml:"length"`
}

func main(){
	song := Song{Name: "测试",Length:126}
	songType := reflect.TypeOf(song)
	fmt.Println("struct name : ", songType.Name())
	numField := songType.NumField()
	songValue := reflect.ValueOf(song)
	for i := 0; i < numField; i++ {
		//获取结构的字段
		field := songType.Field(i)
		//打印字段类型的名称
		fmt.Println(field.Type.Name())
		//获取字段名称
		fmt.Println("fieldName:", field.Name)
		//获取字段json标签
		fmt.Println("json tag:",field.Tag.Get("json"))
		//获取字段xml标签
		fmt.Println("xml tag:", field.Tag.Get("xml"))
		//获取字段的值
		value := songValue.Field(i)
		//打印字段的值
		fmt.Println("value :", value)
	}
}

// struct name :  Song
// string
// fieldName: Name
// json tag: name
// xml tag: name
// value : 测试
//============================
// int
// fieldName: Length
// json tag: length
// xml tag: length
// value : 126

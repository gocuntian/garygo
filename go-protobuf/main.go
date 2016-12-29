package main

import (
    "log"
    "fmt"
    //辅助库
    "github.com/golang/protobuf/proto"
    //test.pd.go的路径
    "github.com/xingcuntian/go_test/go-protobuf/example"
)

func main(){
    // 创建一个消息 Test
    test := &example.Test{
        // 使用辅助函数设置域的值
        Label: proto.String("hello"),
        Type:  proto.Int32(17),
        Optionalgroup: &example.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }

    // 进行编码
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }

    // 进行解码
    newTest := &example.Test{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }

    // 测试结果
   // if test.GetType() != newTest.GetType() {
        //log.Fatalf("data mismatch %q != %q", test.GetType(), newTest.GetType())
        fmt.Println(1)
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    //}
    fmt.Println(1)
}
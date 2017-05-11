package main

import (
    "fmt"
    "encoding/json"
    "strings"
    "log"
)
//golang json.Unmarshal里使用interface{}解析json字符串时int64位无法精确

func main(){
    UnmarshalInterface()
    UnmarshalInterface1()
}
// 预期输出结果：

// 
// map[d:map[a:87702850d4926407 b:1 c:143168533585583274] e:]
// 真实输出结果：

// map[d:map[a:87702850d4926407 b:1 c:1.4316853358558328e+17] e:]
// map[d:map[a:87702850d4926407 b:1 c:1.4316853358558328e+17] e:]

func UnmarshalInterface(){
    v:=`{"d":{"a":"87702850d4926407","b":1,"c":143168533585583274},"e":""}`

    var d interface{}

    err :=json.Unmarshal([]byte(v),&d)

    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(d)
}

//解决方法：

func UnmarshalInterface1(){
    v:=`{"d":{"a":"87702850d4926407","b":1,"c":143168533585583274},"e":""}`
    var d interface{}
    decoder := json.NewDecoder(strings.NewReader(v))
    decoder.UseNumber()

    if err := decoder.Decode(&d); err != nil {
        log.Fatal(err)
    }

    fmt.Println(d)
}
//map[e: d:map[a:87702850d4926407 b:1 c:143168533585583274]]
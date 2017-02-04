package main
//Go有很多种数据类型，包括字符串类型，整型，浮点型，布尔型等等，这里有几个基础的例子。
import "fmt"

func main(){
     // 字符串可以使用"+"连接
     s:="go"+"lang"
     fmt.Println(s)

     //整型和浮点型
     d:=1+1
     fmt.Println(d)

     f:=7.0/3.0
     fmt.Println(f)
     
     // 布尔型的几种操作符
     fmt.Println(true && false)
     fmt.Println(true || false)
     fmt.Println(!true)


}
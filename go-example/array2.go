package main

import "fmt"

// 数组作为参数时，函数内部不改变数组内部的值，除非是传入数组的指针。
// 数组的指针：*[3]int
// 指针数组：[2]*int

func main(){
    a:=[...]User{
        {0,"user0"},
        {1,"user1"},
    }
    b:=[...]*User{
        {0,"user0"},
        {1,"user1"},
    }
    e:=&[...]User{
        {0,"user0"},
        {1,"user1"},
    }
    fmt.Println(a,len(a))
    fmt.Println(b,len(b))

    
    fmt.Println("=================")
    fmt.Println(e,len(e))
    fmt.Println(*e,len(e))
    fmt.Println("=================\r\n")
    c:= &[5]int{1,2,3,4,5}//数组的指针
    fmt.Println(c)

}

type User struct{
    Id int
    Name string
}

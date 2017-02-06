package main

import "fmt"
// 使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。
// 访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。

type Person struct{
    Id int
    Name string
}

type Student struct{
    Person
    Score int
}

func (this Person) test(){
    fmt.Println("Person test")
}

func (this Student) test(){
    fmt.Println("Student test")
}

func main(){
    p:=Student{Person{2,"张三"},25}
    p.test()
}
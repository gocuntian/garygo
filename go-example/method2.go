package main

// 从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。
// instance.method(args)->(type).func(instance,args)
// 为了区别这两种方式，官方文档中将左边的称为Method Value，右边则是Method Expression。Method Value是包装后的状态对象，总是与特定的对象实例关联在一起（类似闭包，拐带私奔），而Method Expression函数将Receiver作为第一个显式参数，调用时需额外传递。
// 注意：对于Method Expression，T仅拥有T Receiver方法，T拥有（T+T）所有方法。
import "fmt"

type Person struct{
    Id int
    Name string
}

func (this Person) test(x int){
    fmt.Println("Id: ",this.Id,"Name: ",this.Name)
    fmt.Println("x=",x)
}

func (this Person) test2(tag string) string{
    return this.Name + tag
}

func main(){
    p:=Person{2,"张三"}
    p.test(1)

    var f1 func(int) = p.test
    f1(2)
    Person.test(p,3)
    var f2 func(Person,int) = Person.test
    f2(p,4)
    
    
    // var f2 func(string) string = p.test2
    // res:= f2("tag")
    // fmt.Println(res)



}
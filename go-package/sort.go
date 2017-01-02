package main

import (
    "fmt"
    "sort"
)
//sort包提供了排序切片和用户自定义数据集的函数。
type Person struct{
    Name string
    Age int
}

func (p Person) String()string{
    return fmt.Sprintf("%s : %d\n",p.Name,p.Age)
}

type ByAge []Person

func (a ByAge) Len()int{
    return len(a)
}

func (a ByAge) Swap(i,j int){
    a[i],a[j]=a[j],a[i]
}

func (a ByAge) Less(i,j int)bool{
    return a[i].Age < a[j].Age
}

func main(){
    people:=[]Person{
        {"Bob",21},
        {"John",32},
        {"macaron",43},
        {"Jenny",26},
    }
    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)
}
package main

import (
    "fmt"
    "container/list"
)

func main(){
    l:=list.New()//创建一个链表
    e4:=l.PushBack(4)
    e1:=l.PushFront(1)
    l.InsertBefore(3,e4)
    l.InsertAfter(2,e1)
    count:=l.Len()
    fmt.Printf("链表(1)中元素的个数:%d\n",count);

    l2:= list.New()
    l2.PushBack(100)
    l2.PushBack(101)
    l2.PushBack(200)
    l2.PushBack(201)
    l2.PushFront(99)
    count2:=l2.Len()
    fmt.Printf("链表(2)中元素的个数:%d\n",count2);

    l.PushBackList(l2)
    l.PushFrontList(l2)

    fmt.Printf("============================>\r\n")
    for e:=l.Front();e!=nil;e=e.Next(){
        fmt.Println(e.Value)
    }
    fmt.Printf("============================>\r\n")

    l3:=list.New();
    l3.PushBack("z")
    l3.PushFront("A")
    count3:=l3.Len()
    fmt.Printf("链表(2)中元素的个数:%d\n",count3);

    // for e:=l2.Front();e!=nil;e=e.Next(){
    //     l.MoveToFront(e)
    // }

    fmt.Printf("=============end===============>\r\n")
    l.MoveToFront(l.Front())
    //fmt.Println(l.Front().Value)
    l.MoveToBack(l.Front().Next())
    
    l.Remove(l.Front().Next())
    l.Remove(l.Front().Next())
    
    
    l.MoveBefore(l.Front().Next(),l.Front())

    l.MoveAfter(l.Front().Next(),l.Front())
    
    for e:=l.Front();e!=nil;e=e.Next(){
        fmt.Println(e.Value)
    }
    // fmt.Printf("\r\n");
    // for e:=l.Back();e!=nil;e=e.Prev(){
    //     fmt.Println(e.Value)
    // }
}
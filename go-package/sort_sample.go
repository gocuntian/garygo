package main

import (
    "fmt"
    "sort"
)

func main(){
    strs:=[]string{"c","a","e","b"}
    sort.Strings(strs)
    fmt.Println("Strings: ",strs)

    fmt.Printf("================================>\r\n")

    ints:=[]int{3,7,1,4,9,2}
    sort.Ints(ints)
    fmt.Println("Ints: ",ints)
    //IntsAreSorted检查a是否已排序为递增顺序
    s:=sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ",s)
}
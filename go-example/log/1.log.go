package main

import (
    "log"
)
// log包主要提供了3类接口。分别是 “Print 、Panic 、Fatal ”，
// 对每一类接口其提供了3中调用方式，分别是 "Xxxx 、 Xxxxln 、Xxxxf"，
// 基本和fmt中的相关函数类似，
func main(){
    arr := []int{2,3}
    log.Print("Print array ",arr,"\n")
    log.Println("Println array ",arr)
    log.Printf("Printf array with item [%d, %d]\n",arr[0],arr[1])
}
// 2017/03/22 18:47:41 Print array [2 3]
// 2017/03/22 18:47:41 Println array  [2 3]
// 2017/03/22 18:47:41 Printf array with item [2, 3]
package main

import (
    "fmt"
    "os"
)
// func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
// Fscanln类似Fscan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。
func main(){
   dir, _ := os.Getwd()
   file, err := os.Open(dir+"/readme2.md")
   if err != nil {
       panic(err)
   }
   defer file.Close()

   var col1, col2 ,col3 []string
   for{
       var v1, v2, v3 string
       _, err := fmt.Fscanln(file, &v1, &v2, &v3)
       if err != nil {
           break
       }
       col1 = append(col1,v1)
       col2 = append(col2,v2)
       col3 = append(col3,v3)
   }
   fmt.Println(col1)
   fmt.Println(col2)
   fmt.Println(col3)
}
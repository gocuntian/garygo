package main

import (
    "fmt"
    "time"
    "errors"
)
//errors包实现了创建错误值的函数。
//func New(text string) error
type MyError struct{
    When time.Time
    What string
}

func (e MyError) Error() string {
    return fmt.Sprintf("%v : %v",e.When,e.What)
}

func oops() error {
    return MyError{
        time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
         "the file system has gone away",
    }
}


func main(){
   if err := oops(); err!=nil{
       fmt.Println(err)
   }

   //func New(text string) error

   err := errors.New("emit macho dwarf: elf header corrupted\r\n")
   if err != nil {
       fmt.Print(err)
   }

   const name, id = "xingcuntian", 24
   err = fmt.Errorf("user %q (id %d) not found\r\n", name, id)
   if err != nil{
       fmt.Print(err)
   }
}
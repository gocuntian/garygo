package main

import (
    "log"
    "os"
    "runtime/trace"
)

func main(){
   file, err := os.Create("trace.out")
   if err != nil {
       log.Fatal(err.Error())
   }
   defer file.Close()

   //func Start(w io.Writer) error

   err = trace.Start(file)
   if err != nil {
       log.Fatal(err.Error())
   }
   defer trace.Stop()
}
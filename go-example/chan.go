package main

import (
    "fmt"
    "runtime"
    "time"
)

func get_sum_of_divisible(num int,divider int,resultChan chan int){
   sum:=0
   for value:=0;value < num;value++ {
       if value%divider == 0{
           sum +=value 
       }
   }
   resultChan <- sum
}

func main(){
    runtime.GOMAXPROCS(runtime.NumCPU())
    LIMIT:=1000
    //用于被１５除结果
    job1:=make(chan int,1)
    //用于被３，５除结果
    job2:=make(chan int,2)

    t_start:=time.Now()
    go get_sum_of_divisible(LIMIT,15,job1)
    go get_sum_of_divisible(LIMIT,3,job2)
    go get_sum_of_divisible(LIMIT,5,job2)

    sum15:=<-job1
    sum3,sum5:=<-job2,<-job2

    sum:=sum3 + sum5 - sum15
    t_end:=time.Now()
    fmt.Println(sum)
    fmt.Println(t_end.Sub(t_start))
}
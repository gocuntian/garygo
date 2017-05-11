package main

import (
    "fmt"
    "runtime"
    "time"
)

func test(c chan bool, n int){
    x:=0
    for i := 0; i < 1000000000; i++ {
        x +=i
    }

    fmt.Println(n,x)
    c <- true
}

func main(){
    	runtime.GOMAXPROCS(1) //设置cpu的核的数量，从而实现高并发
    //runtime.GOMAXPROCS(runtime.NumCPU())

    c := make(chan bool)

    t1 := time.Now().Unix()

    for i :=0 ; i <= 9; i++ {
        go test(c,i)
    }

    var cnt int = 0

label:
    for {
        select {
            case <-c:
                if cnt == 9 {
                    break label
                }
                cnt += 1
            default:    
        }
    }
    fmt.Println("main ok")
    t2 := time.Now().Unix()

    fmt.Println("use ", (t2 - t1), " s")
}

// 以上代码是我在本地测试多核计算与单核计算的代码，使用单核心时计算 10 次 10亿次加法运算，需要 3s，使用多核心计算时仅需要至多 1s 明显可以看出Golang在多核CPU上做运算时的优势。
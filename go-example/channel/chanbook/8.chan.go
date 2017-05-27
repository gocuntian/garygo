package main

import "fmt"

func main(){
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan) //关闭
	syncChan := make(chan struct{}, 1)
	go func(){
		Loop: //便签
		for {
			select{
				case e, ok := <-intChan: //ok false 通道已关闭且元素已接收完
				      if !ok {
						  fmt.Println("End.")
						  break Loop
					  }
					  fmt.Printf("Received: %v\n",e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}
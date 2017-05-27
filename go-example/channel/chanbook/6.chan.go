package main

import "fmt"

func main(){
	dataChan := make(chan int,5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func(){//接收
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n",elem)
			}else{
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()
	go func() {//发送
		for i:=0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Send: %d [sender]\n",i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}
// Send: 0 [sender]
// Send: 1 [sender]
// Send: 2 [sender]
// Send: 3 [sender]
// Send: 4 [sender]
// Done. [sender]
// Received: 0 [receiver]
// Received: 1 [receiver]
// Received: 2 [receiver]
// Received: 3 [receiver]
// Received: 4 [receiver]
// Done. [receiver]
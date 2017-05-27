package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string,3)
func main(){
	syncChan1 := make(chan struct{},1)
	syncChan2 := make(chan struct{},2)
	go func(){ //用于接收操作
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second...[receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem,"[receiver]")
			}else{
				break
			}
		}
		fmt.Println("Stop. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func(){//用于发送操作
		for _, elem := range []string{"a","b","c","d"}{
			strChan <- elem
			fmt.Println("Sent:",elem,"[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Send a sync signal. [sender]")
			}
		}
		fmt.Println("wait 2 seconds...[sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

// Sent: a [sender]
// Sent: b [sender]
// Sent: c [sender]
// Send a sync signal. [sender]
// Received a sync signal and wait a second...[receiver]
// Sent: d [sender]
// Received: a [receiver]
// wait 2 seconds...[sender]

// Received: b [receiver]
// Received: c [receiver]
// Received: d [receiver]
// Stop. [receiver]

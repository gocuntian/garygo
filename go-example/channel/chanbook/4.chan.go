package main

import (
	"fmt"
	"time"
)
// Counter 代表计数器的类型。
type Counter struct {
	count int
}
var mapChan = make(chan map[string]Counter,1)

func main(){
	syncChan := make(chan struct{}, 2)
	go func(){ //接收
		for {
			if elem, ok := <-mapChan; ok {
				fmt.Println("received val:", elem)
				counter := elem["count"]
				fmt.Println("received counter:", counter)
				counter.count++
				fmt.Println("received count:", counter.count)
			}else{
				break
			}
		}
		fmt.Println("Stopped .[receiver]")
		syncChan <- struct{}{}
	}()

	go func(){ //发送
		countMap := map[string]Counter{
			"count":Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v.[sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

// received val: map[count:{0}]
// received counter: {0}
// The count map: map[count:{0}].[sender]
// received val: map[count:{0}]
// received counter: {0}
// The count map: map[count:{0}].[sender]
// received val: map[count:{0}]
// received counter: {0}
// The count map: map[count:{0}].[sender]
// received val: map[count:{0}]
// received counter: {0}
// The count map: map[count:{0}].[sender]
// received val: map[count:{0}]
// received counter: {0}
// The count map: map[count:{0}].[sender]
// Stopped .[receiver]

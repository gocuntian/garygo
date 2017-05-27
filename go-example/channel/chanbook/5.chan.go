package main

import (
	"fmt"
	"time"
)

// Counter 代表计数器的类型。
type Counter struct {
	count int
}

func (counter *Counter) String() string {
	return fmt.Sprintf("{count: %d}", counter.count)
}

var mapChan = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() { //接收
		for {
			if elem, ok := <-mapChan; ok {
				fmt.Println("received val:", elem)
				counter := elem["count"]
				fmt.Println("received counter:", counter)
				counter.count++
				fmt.Println("received count:", counter.count)
			} else {
				break
			}
		}
		fmt.Println("Stopped .[receiver]")
		syncChan <- struct{}{}
	}()

	go func() { //发送
		countMap := map[string]*Counter{
			"count": &Counter{},
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

//指针
// received val: map[count:{count: 0}]
// received counter: {count: 0}
// received count: 1
// The count map: map[count:{count: 1}].[sender]
// received val: map[count:{count: 1}]
// received counter: {count: 1}
// received count: 2
// The count map: map[count:{count: 2}].[sender]
// received val: map[count:{count: 2}]
// received counter: {count: 2}
// received count: 3
// The count map: map[count:{count: 3}].[sender]
// received val: map[count:{count: 3}]
// received counter: {count: 3}
// received count: 4
// The count map: map[count:{count: 4}].[sender]
// received val: map[count:{count: 4}]
// received counter: {count: 4}
// received count: 5
// The count map: map[count:{count: 5}].[sender]
// Stopped .[receiver]

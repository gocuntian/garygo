package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main(){
	syncChan := make(chan struct{}, 2)
	go func(){//用于接收数据
		for{
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			}else{
				break
			}
		}
		fmt.Println("Stopped. [receiver]\n")
		syncChan <- struct{}{}
	}()

	go func(){//用于发送数据
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			fmt.Printf("Before send :%v.[sender]\n",countMap)
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map :%v.[sender]\n",countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
// map 类型引用类型 且 是通道类型的元素值类型
// 接收方对 mapChan 改变会影响源 mapChan 
// The count map :map[count:1].[sender]
// The count map :map[count:2].[sender]
// The count map :map[count:3].[sender]
// The count map :map[count:4].[sender]
// The count map :map[count:5].[sender]
// Stopped. [receiver]


// Before send :map[].[sender]
// The count map :map[count:1].[sender]
// Before send :map[count:1].[sender]
// The count map :map[count:2].[sender]
// Before send :map[count:2].[sender]
// The count map :map[count:3].[sender]
// Before send :map[count:3].[sender]
// The count map :map[count:4].[sender]
// Before send :map[count:4].[sender]
// The count map :map[count:5].[sender]
// Stopped. [receiver]

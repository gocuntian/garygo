package main

import (
	"fmt"
	"time"
)

func main(){
	timer := time.NewTimer(2*time.Second)
	fmt.Printf("Present time:%v.\n",time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time:%v.\n",expirationTime)
	fmt.Printf("Stop timer:%v.\n",timer.Stop())
}

// Present time:2017-05-17 09:59:11.618733145 +0800 CST.
// Expiration time:2017-05-17 09:59:13.618850429 +0800 CST.
// Stop timer:false.

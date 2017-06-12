package slab

import (
	"fmt"
	"testing"
)

func Test_ChanPool_AllocAndFree(t *testing.T) {
	//_ = NewChanPool(2, 1*10, 2, 2*10)

	page := make([]byte, 20)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	m := page[i*2 : (i+1)*2 : (i+1)*2]
	// 	fmt.Println(m)
	// }
	s1 := page[18:20]
	s2 := page[18:20:20]
	s3 := page[20:20]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// fmt.Println(page)

}

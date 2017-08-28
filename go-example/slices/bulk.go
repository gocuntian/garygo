package main

import (
	"fmt"
)

type Value []interface{}

func main() {
	batch := make([]Value, 0)
	fmt.Println("Slice batch = ", batch, ", len = ", len(batch), "cap = ", cap(batch)) //Slice batch =  [] , len =  0 cap =  0
	for i := 0; i < 10; i++ {
		batch = append(batch, Value{i, 1.5, "xingcuntian"})
	}

	for _, v := range batch {
		fmt.Println(v)
	}
}

// [0 1.5 xingcuntian]
// [1 1.5 xingcuntian]
// [2 1.5 xingcuntian]
// [3 1.5 xingcuntian]
// [4 1.5 xingcuntian]
// [5 1.5 xingcuntian]
// [6 1.5 xingcuntian]
// [7 1.5 xingcuntian]
// [8 1.5 xingcuntian]
// [9 1.5 xingcuntian]

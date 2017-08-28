package main

import (
	"fmt"
	"sort"
)

func main() {
	chapts := make(map[int]string)
	chapts[1] = "this is one"
	chapts[3] = "this is  three"
	chapts[2] = "this is two"
	

	for k, v := range chapts {
		fmt.Println(k ,"=>",v)
	}

	var keys []int
	for k := range chapts {
		keys = append(keys,k)
	}

	sort.Ints(keys)
	fmt.Println("After sorting")
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", chapts[k])
	}

}
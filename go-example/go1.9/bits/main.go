package main

import (
	"fmt"
	"math/bits"
)

func main() {
	const n = 100
	//start omit
	fmt.Printf("%d  (%b) has %d bits set one\n", n, n, bits.OnesCount(n))
	//100  (1100100) has 3 bits set one
	fmt.Printf("%d reversed is %d\n", n, bits.Reverse(n))
	//100 reversed is 2738188573441261568
	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))
	//100 can be encoded in 7 bits
	// end omit
}

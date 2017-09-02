package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlockK("Send 1 BTC to gary1")
	bc.AddBlockK("Send 2 more to gary2")
	bc.AddBlockK("Send 3 to gary3")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

// Data: Gary Block
// Hash: 335ff6f2ba6f966b87c5456dbaf51dc4deae5f18778680021271b7baa5eb3729

// Prev. hash: 335ff6f2ba6f966b87c5456dbaf51dc4deae5f18778680021271b7baa5eb3729
// Data: Send 1 BTC to gary1
// Hash: 4fbf65930abfcb4f0196e4370c1bfe79739c04118d10a042457ed53714bbf6f8

// Prev. hash: 4fbf65930abfcb4f0196e4370c1bfe79739c04118d10a042457ed53714bbf6f8
// Data: Send 2 more to gary2
// Hash: d865385fd04d8aa8c4a5fa6bf1edd8ac7ae9332d4b3bfd690dfe8e667eef2386

// Prev. hash: d865385fd04d8aa8c4a5fa6bf1edd8ac7ae9332d4b3bfd690dfe8e667eef2386
// Data: Send 3 to gary3
// Hash: 1a61b55788b2eac730f84f4a10d9cb3021d09cde15ae0d8eb49d8bd4a7e5d45f

// [[] [71 97 114 121 32 66 108 111 99 107] [49 53 48 52 50 53 52 56 51 57]]

// ============================================================================ []

// [71 97 114 121 32 66 108 111 99 107 49 53 48 52 50 53 52 56 51 57]

// [
//	[213 214 65 140 159 159 228 228 209 200 133 11 121 149 240 121 39 228 68 34 28 10 19 134 225 34 191 207 153 102 176 44]
// 	[83 101 110 100 32 49 32 66 84 67 32 116 111 32 56 51 57]
//  ]
// ============================================================================ []

// [
// 213 214 65 140 159 159 228 228 209 200 133 11 121 149 240 121 39 228 68 34 28 10 19 134 225 34 191 207 153 102 176 44
// 83 101 110 100 32 49 32 66 84 67 32 116 111 32 1 51 57
//]

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	fmt.Println([][]byte{b.PrevBlockHash, b.Data, timestamp})
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	fmt.Println("============================================================================", []byte{}, "\r\n")
	fmt.Println(headers)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
	}
	block.SetHash()
	return block
}

func NewGaryBlock() *Block {
	return NewBlock("Gary Block", []byte{})
}

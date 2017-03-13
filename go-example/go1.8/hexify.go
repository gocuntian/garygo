package main

import "encoding/hex"

func Hexify(in string) string {
	return hex.EncodeToString([]byte(in))
}
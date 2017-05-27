package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	server := "127.0.0.1:5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal err : %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for i := 0; i < 50; i++ {
		msg := RandString(i)
		msgLen := fmt.Sprintf("%03s", strconv.Itoa(len(msg)))
		words := "aaaa" + msgLen + msg
		fmt.Println(len(words), words)
		conn.Write([]byte(words))
	}
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(3)
		if t == 0 {
			rs = append(rs, strconv.Itoa(rand.Intn(10)))
		} else if t == 1 {
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

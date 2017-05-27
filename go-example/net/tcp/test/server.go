package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func main() {
	netListen, _ := net.Listen("tcp", ":5000")
	defer netListen.Close()
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	allbuf := make([]byte, 0)
	buffer := make([]byte, 1024)

	for {
		readLen, err := conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read error")
			return
		}
		if len(allbuf) != 0 {
			allbuf = append(allbuf, buffer...)
		} else {
			allbuf = buffer[:]
		}
		var readP int = 0
		for {
			if readLen-readP < 7 {
				allbuf = buffer[readP:]
				break
			}
			msgLen, _ := strconv.Atoi(string(allbuf[readP+4 : readP+7]))
			logLen := 7 + msgLen
			if len(allbuf[readP:]) >= logLen {
				fmt.Println(string(allbuf[readP : readP+logLen]))
				readP += logLen
				if readP == readLen {
					allbuf = nil
					break
				}
			} else {
				allbuf = buffer[readP:]
				break
			}
		}
	}
}

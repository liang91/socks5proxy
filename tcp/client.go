package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

var size = 4000
var recvLen = 1024 * 1024

func main() {
	hosts := []string{
		"10.189.52.177:9001",
	}
	for _, host := range hosts {
		for count := 1; count <= size; count++ {
			conn, err := net.Dial("tcp", host)
			if err != nil {
				fmt.Println(count, "Dial ERROR", err)
				break
			}
			fmt.Println(count, host, "DIAL SUCCESS")
			go communicate(conn)
		}
	}

	time.Sleep(time.Duration(1) * time.Hour)
}

func communicate(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buf := make([]byte, recvLen)
	for {
		// send content request
		_, err := conn.Write([]byte("req content\n"))
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Server Close, Gotta Close Client, ERROR:", err)
			} else {
				fmt.Println("Server Occurs Exception, ERROR:", err)
			}
			return
		}

		// read content
		_, err = reader.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Server Close, Gotta Close Client, ERROR:", err)
			} else {
				fmt.Println("Server Occurs Exception, ERROR:", err)
			}
			return
		}

		time.Sleep(3 * time.Second)
	}
}

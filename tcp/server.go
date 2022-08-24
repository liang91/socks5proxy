package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var addr = "10.189.52.177:9001"

var sendLen = 1024 * 1024

func main() {
	listener, _ := net.Listen("tcp", addr)
	var count int
	for {
		conn, err := listener.Accept()
		count++
		if err != nil {
			fmt.Println(count, "Accept ERROR:", err)
			break
		}
		fmt.Println(count, "Client Accept:", conn.RemoteAddr())
		go cdn(conn)
	}

	time.Sleep(time.Hour)
}

func cdn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	content := make([]byte, 0, sendLen)
	for len(content) < cap(content) {
		content = append(content, '1')
	}

	for {
		req, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("Client closed")
			} else {
				fmt.Println("Unknown Client ERROR", err)
			}
			return
		}

		if strings.Trim(req, "\n") == "req content" {
			_, err = conn.Write(content)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("Client closed")
				} else {
					fmt.Println("Unknown Client ERROR", err)
				}
				return
			}
		}
	}
}

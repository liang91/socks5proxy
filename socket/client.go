package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("hello world"))
	conn.Close()
}

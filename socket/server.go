package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	buf := make([]byte, 1024)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		nr, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			if errors.Is(err, io.EOF) {
				fmt.Println("Client Close")
			}
		}
		log.Println(string(buf[:nr]))

		time.Sleep(time.Second)

		_, err = conn.Write(buf[:nr])
		if err != nil {
			fmt.Println(err)
		}
		conn.Close()
	}
}

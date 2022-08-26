package main

import (
	"bytes"
	"log"
	"net"
)

func handleProxy(conn net.Conn) {
	defer conn.Close()

	remote := conn.RemoteAddr()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println(remote, "Read ERROR:", err)
			break
		}
		log.Println(remote, " Receive:", string(buf[:n]))

		if bytes.Contains(buf, []byte("Bye")) {
			break
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Println(remote, "Write ERROR:", err)
			break
		}
	}
	log.Println(remote, " Close")
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleProxy(conn)
	}
}

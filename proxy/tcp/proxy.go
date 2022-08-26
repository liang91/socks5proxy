package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleClient(client)
	}
}

func handleClient(client net.Conn) {
	defer client.Close()

	server, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		_, err = io.Copy(server, client)
		if err != nil {
			log.Println("Copy Client to Server Error:", err)
		} else {
			log.Println("Copy Client to Server Complete")
		}
		server.Write([]byte("Bye"))
		wg.Done()
	}()

	go func() {
		_, err = io.Copy(client, server)
		if err != nil {
			log.Println("Copy Server to Client Error:", err)
		} else {
			log.Println("Copy Server to Client Complete")
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Proxy Exit")
}

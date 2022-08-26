package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	for {
		client, err := listener.Accept()
		if err != nil || client == nil {
			log.Println(err)
			continue
		}

		go func() {
			defer client.Close()
			buf := make([]byte, 1024)
			for {
				_, err := client.Read(buf)
				if err != nil {
					log.Println(err)
					return
				}
				if bytes.Contains(buf, []byte("\r\n\r\n")) {
					break
				}
			}
			end := bytes.LastIndex(buf, []byte{'\r', '\n', '\r', '\n'})
			content := string(buf[:end])
			fmt.Println(content)
			fmt.Println("================================================")
			lines := strings.Split(content, "\r\n")
			parts := strings.Split(lines[0], " ")
			if parts[0] == "CONNECT" {
				server, err := net.Dial("tcp", parts[1])
				if err != nil {
					log.Println(err)
					return
				}
				fmt.Fprint(client, "HTTP/1.1 200 Connection Established\r\n\r\n")
				go io.Copy(server, client)
				io.Copy(client, server)
			}
		}()
	}
}

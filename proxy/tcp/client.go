package main

import (
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var COUNT = 1

func main() {
	var wg sync.WaitGroup
	wg.Add(COUNT)

	for count := 1; count <= COUNT; count++ {
		go func() {
			defer wg.Done()

			proxy, err := net.Dial("tcp", "127.0.0.1:3000")
			if err != nil {
				log.Fatal(err)
			}
			defer proxy.Close()

			data := []byte("hello world ")
			recv := make([]byte, 1024)
			for i := 1; i <= 10000; i++ {
				data := strconv.AppendInt(data, int64(i), 10)
				_, err := proxy.Write(data)
				if err != nil {
					log.Println("Client Write Failed:", err)
					return
				}

				n, err := proxy.Read(recv)
				if err != nil {
					log.Print("Client Read Failed:", err)
					return
				}
				log.Println("Recv", string(recv[:n]))

				time.Sleep(time.Millisecond * 10)
			}
			proxy.Write([]byte("Bye"))
			log.Println("Client Exit")
		}()
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	var size = 10000
	resps := make([]*http.Response, 0, size)
	for m := 1; m <= 1000; m++ {
		go func() {
			for i := 1; i <= size; i++ {
				resp, err := http.Get("http://localhost:8090/hello")
				if err != nil {
					log.Fatal(err)
				}
				resps = append(resps, resp)

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print(i, " ", string(body))
			}
		}()

	}

	time.Sleep(time.Minute)

	for _, resp := range resps {
		resp.Body.Close()
	}
}

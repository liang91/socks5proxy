package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://mojotv.cn/assets/image/logo01.png")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dst, err := os.Create("xingxing.png")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

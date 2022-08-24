package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go func() {
		mux1 := http.NewServeMux()
		mux1.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.RemoteAddr)
			fmt.Fprintln(w, "are you ok?")
		})
		http.ListenAndServe(":8080", mux1)
	}()

	go func() {
		mux2 := http.NewServeMux()
		mux2.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.RemoteAddr)
			fmt.Fprintln(w, "hello world!")
		})
		http.ListenAndServe(":9090", mux2)
	}()

	time.Sleep(time.Hour)
}

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Query().Get("name"))
	fmt.Fprintln(w, "hello from server")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for key, vals := range req.Header {
		for _, val := range vals {
			fmt.Fprintf(w, "%v: %v\n", key, val)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", headers)

	http.ListenAndServe(":8090", nil)
}

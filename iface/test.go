package main

import (
	"fmt"
	"net/http"
)

func main() {
	items := make([]interface{}, 0, 3)
	items = append(items, 0)
	items = append(items, "are you ok?")
	items = append(items, http.NewServeMux())

	for _, item := range items {
		if val, ok := item.(int); ok {
			fmt.Println(val)
		} else if val, ok := item.(string); ok {
			fmt.Println(val)
		} else if val, ok := item.(*http.ServeMux); ok {
			fmt.Println(val)
		}
	}
}

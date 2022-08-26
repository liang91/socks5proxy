package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123456000"
	right := strings.TrimRight(str, "0")
	fmt.Println(right)
}

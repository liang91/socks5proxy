package main

import "fmt"

func main() {
	sli := make([]byte, 4)
	sli[0] = 'a'
	sli[1] = 'b'
	sli[2] = 'c'
	sli[3] = 'd'
	fmt.Println(string(sli))
	copy(sli, []byte{'1', '2', '3'})
	fmt.Println(string(sli))
}

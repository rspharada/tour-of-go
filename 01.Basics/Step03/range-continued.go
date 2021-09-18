package main

import "fmt"

func main() {
	pow := make([]int, 10) // len=10 cap=10 [0 0 .... 0]
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for key, value := range pow {
		fmt.Printf("%d %d\n", key, value)
	}
}

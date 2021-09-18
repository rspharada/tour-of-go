package main

import "fmt"

func test() bool {
	defer fmt.Println("world")

	fmt.Println("hello")

	return true
}
func main() {
	fmt.Println(test())
}

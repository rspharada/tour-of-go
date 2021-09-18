package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// low以上 hight未満
	var s []int = primes[1:4]
	fmt.Println(s, len(s), cap(s))

	var s1 []string
	fmt.Println(s1, len(s1), cap(s1))

	s2 := [...]string{"A", "B"}
	fmt.Println(s2, len(s2), cap(s2))
}

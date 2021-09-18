package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s1 := s[1:4]
	fmt.Println(s1) // 3 5 7

	s2 := s[:2]
	fmt.Println(s2) // 2 3

	s3 := s[1:]
	fmt.Println(s3) // 3 5 7 11 13

	s = s[1:4]
	fmt.Println(s) // 3 5 7

	s = s[:2]
	fmt.Println(s) // 3 5

	s = s[1:]
	fmt.Println(s, len(s), cap(s)) // [5] 1 4

	s = s[0:4]
	fmt.Println(s, len(s), cap(s)) // [5 7 11 13] 4 4
}

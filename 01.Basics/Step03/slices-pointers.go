package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]   // "John" "Paul"
	b := names[1:3]   // "Paul" "George"
	fmt.Println(a, b) // "John" "Paul" "George" "Ringo"

	b[0] = "XXX"
	fmt.Println(a, b)
	// a = "John" "XXX"
	// b = "XXX" "George"
	fmt.Println(names)
	// "John" "XXX" "George" "Ringo"
}

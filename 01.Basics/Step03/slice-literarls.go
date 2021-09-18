package main

import "fmt"

func main() {
	// {2, 3, 4, 7, 11, 13}という配列（基底配列）を作成し、それを参照するスライスを宣言します。
	q := []int{2, 3, 4, 7, 11, 13}
	fmt.Println(q, len(q), cap(q))

	// {true, false, true, true, false, true}という配列（基底配列）を作成し、それを参照するスライスを宣言します。
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r, len(r), cap(r))

	// 上記と同様
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s, len(s), cap(s))
}

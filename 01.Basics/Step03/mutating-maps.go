package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 要素の削除
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// キーに対する要素が存在するかのチェック
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

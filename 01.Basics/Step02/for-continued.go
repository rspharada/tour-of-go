package main

import "fmt"

func main() {
	sum := 1
	// 初期化と後処理ステートメントの記述は省略できる
	// goにはwhileやdo-whileが存在しない
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

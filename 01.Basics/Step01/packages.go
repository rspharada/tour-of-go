package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	* 数を強制的に変える場合は、乱数生成でシードを与える必要があります。rand.Seedを見てみてください。
	* playground 上での時間は一定なので他のものをシードとして使う必要があります。)
	 */
	rand.Seed(time.Now().UnixNano()) // seedがないと常に同じ値が出力される
	fmt.Println("My favorite number is", rand.Intn(10))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// Goでは、最初の文字が大文字で始まる名前は、外部のパッケージから参照（エクスポート：公開）できるようになる
	// mathパッケージのPiはエクスポートされているもの。Piをpiにして実行するとエラーになる
	fmt.Println(math.Pi)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// goのswitchはcaseの最後にbreakが暗黙で実行される。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
		fallthrough // breakを実行したくなければ、fallthroughを仕込む
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

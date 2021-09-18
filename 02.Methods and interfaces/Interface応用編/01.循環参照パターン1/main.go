package main

import (
	"fmt"
	"go-sample/a"
)

func main() {
	a := a.NewA()
	fmt.Println(a.Get())
}

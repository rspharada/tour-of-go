package a

import (
	"go-sample/b"
)

type A struct{}

func NewA() *A {
	return new(A)
}

func (a *A) Get() string {
	b := &b.B{A2: NewA2()}
	return b.Get()
}

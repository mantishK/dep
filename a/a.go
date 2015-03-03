package a

import (
	"fmt"

	"github.com/mantishK/dep/b"
)

type A struct {
}

func (a A) PrintA() {
	fmt.Println(a)
}

func NewA() *A {
	a := new(A)
	return a
}

func RequireB() {
	o := b.NewB()
	o.PrintB()
}

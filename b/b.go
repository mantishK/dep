package b

import (
	"fmt"

	"github.com/mantishK/dep/i"
)

type B struct {
}

func (b B) PrintB() {
	fmt.Println(b)
}

func NewB() *B {
	b := new(B)
	return b
}

func RequireA(o i.Aprinter) {
	o.PrintA()
}

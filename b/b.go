package b

import (
	"fmt"

	"github.com/mantishK/dep/a"
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

func RequireA() {
	o := a.NewA()
	o.PrintA()
}

package c

import (
	"github.com/mantishK/dep/a"
	"github.com/mantishK/dep/b"
)

func PrintC() {
	o := a.NewA()
	b.RequireA(o)
}

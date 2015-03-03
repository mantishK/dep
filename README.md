# dep
Dealing with dependency cycles in Golang

Go doesn't allow import cycles to occur. If there are any import cycles detected, it throws a compile time error. Generally import cycles are considered as a bad design.    

Import cycles are caused when a package 'a' depends on 'b' and 'b' in turn depends on 'a'. Following Go code illustrates the classic problem of import cycle, AKA dependency cycle.
```Go
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

```

```Go
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

```

This results in the following error during compile time:

```
import cycle not allowed
package github.com/mantishK/dep/a
  imports github.com/mantishK/dep/b
  imports github.com/mantishK/dep/a
```

### How do we avoid this?

The problem:
```
A depends on B 
B depends on A
```
In order to avoid the cyclic dependency, we must introduce an interface in a new package say i. This interface will have all the methods that are in struct A and are accessed by struct B.    
```Go
package i

type Aprinter interface {
  PrintA()
}

```

By doing so, we can make the package b to import i rather than a. Now, package b looks like this:

```Go
package b

import (
  "fmt"

  "github.com/mantishK/dep/i"
)


func RequireA(o i.Aprinter) {
  o.PrintA()
}
```

This still doesn't solve the problem completely, we still need an instance of struct A in the methods of struct B. To solve this we need another package, say c. Package c will import both a and b, it creates an instance of a and passes it to b. 

```Go
package c

import (
  "github.com/mantishK/dep/a"
  "github.com/mantishK/dep/b"
)

func PrintC() {
  o := a.NewA()
  b.RequireA(o)
}
```

The cycle is broken, yet the functionality is achieved.
```
A depends on B
B depends on I
C depends on A and B
```
Problem solved !!    

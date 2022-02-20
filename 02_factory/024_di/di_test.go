package di_test

import (
	"fmt"
	di "github.com/1ikc/design-pattern/02_factory/024_di"
)

type A struct {
	B *B
}
func NewA(b *B) *A {
	return &A{B: b}
}

type B struct {
	C *C
}
func NewB(c *C) *B {
	return &B{C: c}
}

type C struct {
	Num int
}
func NewC() *C {
	return &C{Num: 1}
}

func Example_di() {
	container := di.New()
	if err := container.Provide(NewA); err != nil {
		panic(err)
	}
	if err := container.Provide(NewB); err != nil {
		panic(err)
	}
	if err := container.Provide(NewC); err != nil {
		panic(err)
	}

	err := container.Invoke(func(a *A) {
		fmt.Printf("%d", a.B.C.Num)
	})
	if err != nil {
		panic(err)
	}

	// Output:
	// 1
}
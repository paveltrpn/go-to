package main

import "fmt"

type fracted interface {
	float32 | float64
}

type SFoo[T fracted] struct {
	Name  string
	Field T
}

func sum[T fracted](a, b SFoo[T]) SFoo[T] {
	return SFoo[T]{a.Name + b.Name, a.Field + b.Field}
}

func newSFoo[T fracted](name string, field T) SFoo[T] {
	return SFoo[T]{name, field}
}

func main() {
	var (
		a_var SFoo[float32]
	)

	a_var = newSFoo[float32]("a_var", 100)
	b_var := newSFoo[float32]("b_var", 200)

	sm := sum(a_var, b_var)

	fmt.Printf("sm.Name = %v; sm.Field = %v\n", sm.Name, sm.Field)
}

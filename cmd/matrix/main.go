package main

import (
	"go-to/intern/matrix"
)

func main() {
	foo := matrix.MakeEmpty[float32](4, 4)
	foo.SetIdtt()

	bar := matrix.MakeFrom(*foo)

	baz := matrix.MakeEmpty[float32](4, 4)
	baz.CopyFrom(*foo)

	foo.Print()
	bar.Print()
	baz.Print()
}

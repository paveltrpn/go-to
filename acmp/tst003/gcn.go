package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	a := 2205
	b := 2475

	fmt.Printf("GCN for %v and %v is:\n", a, b)
	fmt.Println(GCD(a, b))
}

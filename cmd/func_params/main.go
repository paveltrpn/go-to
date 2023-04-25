package main

import "fmt"

type matrix struct {
	rows     int
	columnes int

	data []float32
}

type MatrixOption func(*matrix)

func WithRows(rows int) MatrixOption {
	return func(rt *matrix) {
		rt.rows = rows
	}
}

func WithColumnes(cols int) MatrixOption {
	return func(rt *matrix) {
		rt.columnes = cols
	}
}

func MakeMatrix(params ...MatrixOption) *matrix {
	const (
		defRows int = 1
		defCols int = 1
	)

	var rt matrix

	rt.rows = defRows
	rt.columnes = defCols

	for _, p := range params {
		p(&rt)
	}

	rt.data = make([]float32, rt.rows*rt.columnes)

	return &rt
}

func main() {
	mtrx := MakeMatrix(WithColumnes(100))

	fmt.Printf("rows = %v, cols = %v\n", mtrx.rows, mtrx.columnes)
}

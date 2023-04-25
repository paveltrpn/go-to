package matrix

import "fmt"

type Matrix[T float32 | float64] struct {
	columnes int
	rows     int

	data []T
}

func MakeEmpty[T float32 | float64](rows int, columnes int) *Matrix[T] {
	var (
		rt Matrix[T]
	)

	rt.rows = rows
	rt.columnes = columnes
	rt.data = make([]T, rows*columnes)

	return &rt
}

func MakeFrom[T float32 | float64](rhs Matrix[T]) *Matrix[T] {
	rt := MakeEmpty[T](rhs.rows, rhs.columnes)
	copy(rt.data, rhs.data)

	return rt
}

func (m *Matrix[T]) CopyFrom(rhs Matrix[T]) {
	// equal size matrices
	if (m.rows == rhs.rows) && (m.columnes == rhs.columnes) {
		copy(m.data, rhs.data)
		return
	} else {
		nd := make([]T, rhs.rows*rhs.columnes)

		m.rows = rhs.rows
		m.columnes = rhs.columnes
		m.data = nd
	}
}

// Only fo sqaure matrices
func (m *Matrix[T]) SetIdtt() {
	if m.rows != m.columnes {
		return
	}

	for elem := range m.data {
		if (elem % (m.rows + 1)) == 0 {
			m.data[elem] = 1.0
		} else {
			m.data[elem] = 0.0
		}
	}

}

func (m *Matrix[T]) Mult(rhs Matrix[T]) {
	// works only if matrices is mxn and nxp
	if m.columnes != rhs.rows {
		return
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < rhs.columnes; j++ {

		}
	}
}

func (m *Matrix[T]) Invert() {

}

func (m *Matrix[T]) Tranpose() {

}

func (m *Matrix[T]) GetDeterminant() {

}

func (m Matrix[T]) Print() {
	for i, elem := range m.data {
		if i%(m.rows) == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%v ", elem)
	}
	fmt.Printf("\n")
}

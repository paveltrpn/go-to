// Functions below taken from paper - "Parallel programming in Go and Scala"
// Author: Carl Johnell
//
// this code describe sequintial and parallel matrix multiplication
// and matrix chain multiplication algorithms.
//
// This paper contain thesis about how better store a matrix in
// golang - as linear slice or slice of slices (2D slice).
// Paper is suggest to store a matrix as a linear slice because
// of better performance, reason is a less overhead to array bounds check.

package main

import (
	"fmt"
	"math"
)

func seqMatrixMult(m1, m2, res []int, n int) []int {
	m1Rows := len(m1) / n
	for i := 0; i < m1Rows; i++ {
		for k := 0; k < n; k++ {
			for j := 0; j < n; j++ {
				res[n*i+j] += m1[i*n+k] * m2[k*n+j]
			}
		}
	}
	return res
}

func parMatrixMult(vp, n int, m1, m2, res []int) []int {
	chunkSize := n / vp
	start := 0
	end := 0
	ch := make(chan int, vp)
	for i := 0; i < vp; i++ {
		start = end
		end += chunkSize
		if (i + 1) == vp {
			end = n
		}
		startPos := start * n
		endPos := end * n
		m1Chunk := m1[startPos:endPos]
		resChunk := res[startPos:endPos]
		go func(m1, m2, res []int, n int, ch chan int) {
			seqMatrixMult(m1, m2, res, n)
			ch <- 1
		}(m1Chunk, m2, resChunk, n, ch)
	}
	for i := 0; i < vp; i++ {
		<-ch
	}
	return res
}

// Matrix chain mulptiplication section
func computeCost(cost, chain []int, n, i, j int) {
	cost[i*n+j] = math.MaxInt32
	if i == j {
		cost[i*n+j] = 0
	} else {
		for k := i; k <= (j - 1); k++ {
			q := cost[i*n+k] + cost[(k+1)*n+j] + chain[i-1]*chain[k]*chain[j]
			idx := i*n + j
			if q < cost[idx] {
				cost[idx] = q
			}
		}
	}
}
func seqChainMult(cost, chain []int, n int) {
	for i := n; i >= 1; i-- {
		for j := i; j <= n; j++ {
			computeCost(cost, chain, n+1, i, j)
		}
	}
}

func getBounds(n, vp, id int) (int, int) {
	size := n / vp
	low := id * size
	high := low + size
	if (id + 1) == vp {
		high = n
	}
	return low + 1, high
}

func computeBlock(cost, chain []int, n, i, j, il, ih, jl, jh int) {
	if i == j {
		for ii := ih; ii >= il; ii-- {
			for jj := ii; jj <= jh; jj++ {
				computeCost(cost, chain, n, ii, jj)
			}
		}
	} else {
		for ii := ih; ii >= il; ii-- {
			for jj := jl; jj <= jh; jj++ {
				computeCost(cost, chain, n, ii, jj)
			}
		}
	}
}

func partition(rounds []chan int, finish chan int, cost, chain []int, i, n, vp int) {
	j := i
	il, ih := getBounds(n, vp, i)
	for ; j < vp; j++ {
		jl, jh := getBounds(n, vp, j)
		computeBlock(cost, chain, n+1, i, j, il, ih, jl, jh)
		if i > 0 {
			rounds[i-1] <- 1
		}
		if j+1 < vp {
			<-rounds[i]
		}
	}
	if il == 1 {
		finish <- 1
	}
}

func parChainMult(rounds []chan int, finish chan int, cost, chain []int, vp, n int) {
	for i := 0; i < vp; i++ {
		go partition(rounds, finish, cost, chain, i, n, vp)
	}
	<-finish
}

func main() {
	fmt.Printf("Concurrent matrix multiplication test")
}

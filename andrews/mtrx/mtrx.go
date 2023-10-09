package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

func makeRandVec(length int) []float32 {
	rt := make([]float32, length)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		rt[i] = float32(rand.Float64()) * math.MaxFloat32
	}

	return rt
}

func makeRandMtrx(rng int) [][]float32 {
	rt := make([][]float32, rng)
	for i := range rt {
		rt[i] = make([]float32, rng)
		rt[i] = makeRandVec(rng)
	}

	return rt
}

func makeIdttMtrx(rng int) [][]float32 {
	rt := make([][]float32, rng)

	// allocate rows
	for k := range rt {
		rt[k] = make([]float32, rng)
	}

	for i := 0; i < rng; i++ {
		for j := 0; j < rng; j++ {
			if i == j {
				rt[i][j] = 1.0
			} else {
				rt[i][j] = 0.0
			}
		}
	}

	return rt
}

func MtrxOut(mtrx [][]float32, bf *bufio.Writer) {
	for _, fline := range mtrx {
		// convert []float32 to string
		line := strings.Trim(strings.Join(strings.Split(fmt.Sprint(fline), " "), " "), "[]")
		bf.WriteString(line + "\n")
	}

	bf.Flush()
}

func MtrxMultSquare(A, B [][]float32) [][]float32 {
	var rt [][]float32
	return rt
}

func main() {
	foo := makeIdttMtrx(5)

	stdOutBuf := bufio.NewWriter(os.Stdout)
	MtrxOut(foo, stdOutBuf)
}

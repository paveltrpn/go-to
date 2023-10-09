package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INFILE  string = "INPUT.TXT"
	OUTFILE string = "OUTPUT.TXT"
)

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	return slc
}

func collapseNumber(num int) int {
	var cn int = num
	var accum int

	for cn > 10 {
		digits := splitInt(cn)
		accum = 0
		for _, d := range digits {
			accum = accum + d
		}
		cn = accum
	}

	return cn
}

func main() {
	inf, err := os.Open(INFILE)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer inf.Close()

	rd := bufio.NewScanner(inf)

	rd.Scan()
	// trim leading and trailing zeroes
	line := strings.Trim(rd.Text(), "0")

	var (
		left, right int
	)

	outf, err := os.Create(OUTFILE)
	defer outf.Close()

	ln := len([]rune(line))
	if ln == 0 {
		outf.WriteString("NO")
		return
	}

	for i := 1; i < ln; i++ {
		left, _ = strconv.Atoi(line[:i])
		right, _ = strconv.Atoi(line[i:])

		if collapseNumber(left) == collapseNumber(right) {
			fmt.Println("YES")
			outf.WriteString("YES")
			return
		}
	}

	outf.WriteString("NO")
}

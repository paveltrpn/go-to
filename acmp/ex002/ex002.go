//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INFILE  string = "INPUT.TXT"
	OUTFILE string = "OUTPUT.TXT"
)

func main() {
	inf, err := os.Open(INFILE)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer inf.Close()

	rd := bufio.NewScanner(inf)

	rd.Scan()
	N, _ := strconv.Atoi(rd.Text())

	outf, err := os.Create(OUTFILE)
	defer outf.Close()

	if N == 0 {
		outf.WriteString(fmt.Sprint(1))
		return
	}

	if N > 0 {
		// Sum of arithmetic progression from 1 to n is S = (S1 + Sn)*Sn) / 2
		outf.WriteString(fmt.Sprint(((1 + N) * N) / 2))
		return
	}

	if N < 0 {
		var sum int = 0
		for i := N; i < 1; i++ {
			sum = sum + i
		}
		// add 1 because of we calculate from -N towards 0 and 1,
		// we must add positive one at whole negative sum!
		outf.WriteString(fmt.Sprint(sum + 1))
		return
	}
}

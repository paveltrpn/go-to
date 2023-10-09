// ex003
// Vasya and Petya go in school in same grade. Yesterday Petya told
// to Vasya about cunning (tricky) way to give a square of number ended by
// five. Now Vasye can easlly obtain square of two-digit (or even three-digit)
// numbers ended by 5. The way is look like this: to obtain a square of number
// ended by 5 is enought to multiply a number obtained by exclue a last digit
// from initial number by next number and add a "25" at the end of result
// of multiplication. Example - to give a 126^2 is enought to mult 12 by 13
// and add a "25" at end - 12*13 = 156, 156 concatanate "25" answer is 15625.
// Write a program which rises a number ended by five in square to Vasya can
// test thier abileties (knowlege, skills).
// Input
// In single line called input.txt written one natural number, ended by 5
// and not greater than 4*10^5
// Output
// In output file called output.txt out one natural number - A^2 without
// leading zeroes

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INFILE  string = "input.txt"
	OUTFILE string = "output.txt"
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
	A, _ := strconv.Atoi(rd.Text())

	outf, err := os.Create(OUTFILE)
	defer outf.Close()

	outf.WriteString(fmt.Sprint(A * A))
}

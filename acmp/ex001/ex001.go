// Required to add two numbers A and B.
//
// Input
// In single line of input file input.txt written two
// natural numbers seperated by space. Value of numbers
// not greater than 10^9
//
// Output
// In single line of output file output.txt need to write
// a single integer number - sum of two readed numbers A and B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	INFILE  string = "input.txt"
	OUTFILE string = "output.txt"
)

func main() {
	inf, err := os.Open(INFILE)

	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer inf.Close()

	rd := bufio.NewScanner(inf)

	rd.Scan()
	numbers := strings.Split(rd.Text(), " ")

	left, _ := strconv.Atoi(numbers[0])
	right, _ := strconv.Atoi(numbers[1])

	outf, err := os.Create(OUTFILE)
	defer outf.Close()

	outf.WriteString(fmt.Sprint(left + right))
}

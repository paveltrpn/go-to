package main

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	str   = []string{"0", "01", "11", "001", "010", "011", "0001", "0101"}
	strCh = []string{"a", "ab", "bb", "aab", "aba", "abb", "aaab", "abab"}
	bt    = []int{0b0, 0b01, 0b11, 0b001, 0b010, 0b011, 0b0001, 0b0101}
)

func main() {
	sort.Strings(str)
	fmt.Println(str)
	sort.Strings(strCh)
	fmt.Println(strCh)
	sort.Ints(bt)
	for i := 0; i < len(bt); i++ {
		n := int64(bt[i])
		fmt.Print(strconv.FormatInt(n, 2))
		fmt.Print(" ")
	}
	fmt.Println()
}

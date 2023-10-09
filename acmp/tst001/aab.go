package main

import "fmt"

var word = [4]string{"A", "B", "V", "G"}

func main() {
	var num int = 1

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				fmt.Println(word[i], word[j], word[k], num)
				num++
			}
		}
	}
}

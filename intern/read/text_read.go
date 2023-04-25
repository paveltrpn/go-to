package read

import (
	"bufio"
	"os"
)

func TextRead(fname string) string {
	dat, err := os.ReadFile(fname)

	if err != nil {
		println("TextRead(): Error! can't find file - %s", fname)
		panic(err)
	}

	return string(dat)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func TextReadLines(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		println("TextReadLines(): Error! can't find file - %s", fname)
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

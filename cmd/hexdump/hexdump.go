package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

const (
	// String of ASCII symbols (exactly 256 chars), non printable symbols describes as dots.
	printableChars string = "................................ !\"#$%&'()*+,-./0123456789:;<=>?@AB" +
		"CDEFGHIJKLMNOPQRSTUVWXYZ[.]^_`abcdefghijklmnopqrstuvwxyz{|}~........." +
		"....................................................................." +
		"..................................................."
)
const lineHeight = 16
const baLength = 256

// Convert a slice of bytes to string. Every character in returned string
// may by a printable ASCII character or dot symbol.
func bytesToPrintable(line []byte) string {
	var rt strings.Builder

	for _, b := range line {
		rt.WriteString(string(printableChars[int(b)]))
	}

	return rt.String()
}

func hexdumpByHexPkg(arr []byte) {
	for i := 0; i < baLength; i = i + lineHeight {
		chars := bytesToPrintable(arr[i : i+lineHeight])
		// Нужно написать свою функцию вывода байта в hex формате
		// тем более это не сложно - определить значения полубайтов
		// и отобразить на массив hex значений
		hexs := hex.EncodeToString(arr[i : i+lineHeight])
		fmt.Println(hexs, chars)
	}
}

func main() {
	// Оставил тут для примера как определять количество символов в строке
	fmt.Println("Number of characters in printable line -", len([]rune(printableChars)))

	byteArray := make([]byte, baLength)
	rand.Read(byteArray)
	// for i := range byteArray {
	// byteArray[i] = byte(i)
	// }

	hexdumpByHexPkg(byteArray)
}

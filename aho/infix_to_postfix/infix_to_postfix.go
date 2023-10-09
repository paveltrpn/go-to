package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var (
	lookahead int
	inExpr    expressionString_s
	outExpr   expressionString_s
)

func expr() {
	term()
	for {
		if lookahead == '+' {
			match('+')
			term()
			outExpr.appendTermToExprStr('+')
		} else if lookahead == '-' {
			match('-')
			term()
			outExpr.appendTermToExprStr('-')
		} else {
			return
		}
	}
}

func term() {
	if unicode.IsNumber(rune(lookahead)) {
		t := lookahead
		match(lookahead)
		outExpr.appendTermToExprStr(t)
	} else {
		fmt.Printf("syntax error in term()! %v - is not number!\n", string(rune(lookahead)))
		os.Exit(1)
	}
}

func match(t int) {
	if lookahead == t {
		_, lookahead = inExpr.getNextChar()
	} else {
		fmt.Println("syntax error in match()!")
		os.Exit(1)
	}
}

// func readOneRune(rd io.Reader) rune {
// in := bufio.NewReader(rd)
// foo, _, _ := in.ReadRune()
//
// return foo
// }

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("read: ")
	text, _ := reader.ReadString('\n')

	if len([]rune(text)) > 1 {
		inExpr = buildExpressionString(text)
	} else {
		// Use default string
		inExpr = buildExpressionString("8-5+3+1-5")
	}

	outExpr = buildExpressionString("")

	_, lookahead = inExpr.getNextChar()

	expr()
	fmt.Println(outExpr.str)
}

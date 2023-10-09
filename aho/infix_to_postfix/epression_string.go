package main

type expressionString_s struct {
	pos int
	str string
}

type expressionStringInterface interface {
	buildExpressionString(str string) expressionString_s
}

func buildExpressionString(str string) expressionString_s {
	return expressionString_s{0, str}
}

func (es *expressionString_s) appendTermToExprStr(term int) {
	es.str = es.str + string(rune(term))
}

// Return a char from expressionString_s.str at position es.pos
// and increments a expressionString_s.pos.
func (es *expressionString_s) getNextChar() (int, int) {
	if es.pos+1 > len([]rune(es.str)) {
		return -1, int(' ')
	}

	rt := es.str[es.pos]
	es.pos = es.pos + 1
	return es.pos, int(rt)
}

func (es *expressionString_s) loadNewString(new string) {
	es.str = new
	es.pos = 0
}

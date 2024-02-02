package ast

import (
	"strconv"
	"unicode"
)

func Parse(expr string) (tree *Ast) {
	num, expr := ExtractNum(expr)
	tree = newValue(num)
	if len(expr) > 0 && IsOperator(rune(expr[0])) {
		op, expr := Operator(expr[0]), expr[1:]
		tree.Append(op, Parse(expr))
	}
	return tree
}

func ExtractNum(s string) (num float64, rest string) {
	i := 0
	for unicode.IsDigit(rune(s[i])) || s[i] == '.' {
		i++
		if i >= len(s) {
			break
		}
	}
	num, _ = strconv.ParseFloat(s[:i], 64)
	return num, s[i:]
}

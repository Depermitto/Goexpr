package ast

import (
	"Goexpr/lists"
	"strconv"
	"strings"
	"unicode"
)

func Parse(expr string) (tree *Ast) {
	var opStack lists.LinkedList[Operator]
	var num float64

	if expr[0] == '(' {
		num, expr = extractPar(expr)
	} else {
		num, expr = extractNum(expr)
	}

	tree = NewValue(num)
	for len(expr) > 0 {
		c := rune(expr[0])
		if IsOperator(c) {
			opStack.PushTail(Operator(c))
			expr = expr[1:]
		} else if op, err := opStack.PopTail(); err == nil {
			if c == '(' {
				inside, rest := extractPar(expr)
				tree.Append(op, NewValue(inside))
				expr = rest
			} else {
				num, expr = extractNum(expr)
				tree.Append(op, NewValue(num))
			}
		} else {
			return nil
		}
	}

	return tree
}

func extractPar(s string) (num float64, rest string) {
	inside, rest, found := strings.Cut(s[1:], ")")
	if !found {
		return 0, ""
	}
	return Parse(inside).Eval(), rest
}

func extractNum(s string) (num float64, rest string) {
	i := 0
	for ; i < len(s) && (unicode.IsDigit(rune(s[i])) || s[i] == '.'); i++ {
	}
	num, _ = strconv.ParseFloat(s[:i], 64)
	return num, s[i:]
}

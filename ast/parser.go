package ast

import (
	"Goexpr/lists"
	"strconv"
	"strings"
	"unicode"
)

func Parse(expr string) (tree *Ast) {
	if strings.Count(expr, "(") != strings.Count(expr, ")") {
		return nil
	}
	expr = strings.TrimSpace(expr)
	expr = strings.ReplaceAll(expr, " ", "")
	return parse(expr)
}

func parse(expr string) *Ast {
	var opStack lists.LinkedList[Operator]
	var num float64
	if expr[0] == '(' {
		num, expr = extractPar(expr)
	} else {
		num, expr = extractNum(expr)
	}

	tree := NewValue(num)
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
	n := 1
	i := 1
	for ; i < len(s) && n != 0; i++ {
		if s[i] == '(' {
			n++
		} else if s[i] == ')' {
			n--
		}
	}
	inside, rest := s[1:i-1], s[i:]
	return parse(inside).Eval(), rest
}

func extractNum(s string) (num float64, rest string) {
	i := 0
	for ; i < len(s) && (unicode.IsDigit(rune(s[i])) || s[i] == '.'); i++ {
	}
	num, _ = strconv.ParseFloat(s[:i], 64)
	return num, s[i:]
}

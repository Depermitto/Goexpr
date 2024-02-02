package ast

import "math"

type Operator rune

var operators = map[Operator]map[string]any{
	'+': {"priority": 1, "func": func(a float64, b float64) float64 { return a + b }},
	'-': {"priority": 1, "func": func(a float64, b float64) float64 { return a - b }},
	'*': {"priority": 2, "func": func(a float64, b float64) float64 { return a * b }},
	'/': {"priority": 2, "func": func(a float64, b float64) float64 { return a / b }},
	'%': {"priority": 2, "func": func(a float64, b float64) float64 { return float64(int(a) % int(b)) }},
	'^': {"priority": 3, "func": func(a float64, b float64) float64 { return math.Pow(a, b) }},
}

func IsOperator(rune rune) bool {
	for k := range operators {
		if k == Operator(rune) {
			return true
		}
	}
	return false
}

func (o Operator) Eval(a float64, b float64) float64 {
	for k, v := range operators {
		if o == k {
			return v["func"].(func(float64, float64) float64)(a, b)
		}
	}
	return 0
}

func (o Operator) Priority() int {
	for k, v := range operators {
		if o == k {
			return v["priority"].(int)
		}
	}
	return 0
}

package ast

import (
	"fmt"
)

type Ast struct {
	value float64
	left  *Ast
	right *Ast
	op    Operator
}

func newExpression(op Operator, left *Ast, right *Ast) *Ast {
	return &Ast{op: op, left: left, right: right}
}

func newValue(value float64) *Ast {
	return &Ast{value: value}
}

func (a *Ast) IsValue() bool {
	return a.left == nil && a.right == nil
}

func (a *Ast) Eval() float64 {
	if a.IsValue() {
		return a.value
	}

	return a.op.Eval(a.left.Eval(), a.right.Eval())
}

func (a *Ast) Append(op Operator, expr *Ast) {
	if a.IsValue() || op.Priority() <= a.op.Priority() {
		left := *a
		*a = *newExpression(op, &left, expr)
	} else {
		right := *a.right
		a.right = newExpression(op, expr, &right)
	}
}

func (a *Ast) String() (s string) {
	if a.IsValue() {
		return fmt.Sprintf("%v", a.value)
	}

	if a.left != nil {
		s += a.left.String()
	}

	s += fmt.Sprintf("%q", a.op)

	if a.right != nil {
		s += a.right.String()
	}

	return s
}

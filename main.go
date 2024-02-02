package main

import (
	"Goexpr/ast"
	"fmt"
	"log"
)

func main() {
	calc := ast.Parse("(1-10)/(2+1)%2")
	if calc == nil {
		log.Fatalln("incorrect construction")
	}
	fmt.Println(calc)
	fmt.Println(calc.Eval())
}

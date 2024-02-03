package main

import (
	"Goexpr/ast"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("Expr: "); scanner.Scan(); fmt.Print("Expr: ") {
		calc := ast.Parse(scanner.Text())
		if calc == nil {
			log.Println("incorrect construction")
			continue
		}
		fmt.Println(calc.Eval())
	}
}

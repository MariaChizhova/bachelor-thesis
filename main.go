package main

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/parser/ast"
	"fmt"
)

func main() {
	//node := parser.Parse("true")
	node := parser.Parse("1 * 2 + 3")
	fmt.Print(ast.Print(node))

}

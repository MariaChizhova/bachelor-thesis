package main

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/parser/ast"
	"fmt"
)

func main() {
	node := parser.Parse("[1][0]")
	fmt.Print(ast.Print(node))
}

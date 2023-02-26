package main

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/parser/ast"
	"fmt"
)

func main() {
	node := parser.Parse("a")
	fmt.Print(ast.Print(node))
}

package main

import (
	"bachelor-thesis/parser"
	"bachelor-thesis/parser/ast"
	"fmt"
)

func main() {
	node := parser.Parse("a and b or c")
	fmt.Print(ast.Print(node))
}

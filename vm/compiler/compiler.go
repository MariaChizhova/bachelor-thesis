package compiler

import (
	"bachelor-thesis/parser/ast"
	"bachelor-thesis/vm/code"
)

type Compiler struct {
	instructions []code.Instructions
	constants    []interface{}
}

// TODO: remove it?
func concatInstructions(s []code.Instructions) code.Instructions {
	out := code.Instructions{}
	for _, ins := range s {
		out = append(out, ins...)
	}
	return out
}

func Compile(node ast.Node) (program *Program, err error) {
	compiler := &Compiler{}
	compiler.compile(node)
	program = &Program{
		Instructions: concatInstructions(compiler.instructions),
		Constants:    compiler.constants,
	}
	return program, nil
}

func (compiler *Compiler) compile(node ast.Node) {
	switch node.Type() {
	case ast.NodeNumber:
		compiler.NodeNumber(node.(*ast.NumberNode))
	case ast.NodeIdentifier:
		compiler.NodeIdentifier(node.(*ast.IdentifierNode))
	case ast.NodeString:
		compiler.NodeString(node.(*ast.StringNode))
	case ast.NodeBool:
		compiler.NodeBool(node.(*ast.BoolNode))
	case ast.NodeNil:
		compiler.NodeNil(node.(*ast.NilNode))
	case ast.NodeUnary:
		compiler.NodeUnary(node.(*ast.UnaryNode))
	case ast.NodeBinary:
		compiler.NodeBinary(node.(*ast.BinaryNode))
	case ast.NodeFunction:
		compiler.NodeFunction(node.(*ast.FunctionNode))
	case ast.NodeArray:
		compiler.NodeArray(node.(*ast.ArrayNode))
	}
}

func (compiler *Compiler) NodeNumber(node *ast.NumberNode) {
	if node.IsInt == true {
		compiler.emit(code.OpConstant, compiler.addConstant(node.Int64))
	} else if node.IsFloat == true {
		compiler.emit(code.OpConstant, compiler.addConstant(node.Float64))
	}
}

func (compiler *Compiler) NodeIdentifier(node *ast.IdentifierNode) {
	// TODO: implement
}

func (compiler *Compiler) NodeString(node *ast.StringNode) {
	compiler.emit(code.OpConstant, compiler.addConstant(node.Value))
}

func (compiler *Compiler) NodeBool(node *ast.BoolNode) {
	if node.Value {
		compiler.emit(code.OpTrue)
	} else {
		compiler.emit(code.OpFalse)
	}
}

func (compiler *Compiler) NodeNil(node *ast.NilNode) {
	compiler.emit(code.OpNil)
}

func (compiler *Compiler) NodeUnary(node *ast.UnaryNode) {
	compiler.compile(node.Node)
	switch node.Operator {
	case "+":
		// Do nothing

	case "-":
		compiler.emit(code.OpMinus)
	case "not":
		compiler.emit(code.OpNot)
	}
	// compiler.emit(code.OpPop)
}

func (compiler *Compiler) NodeBinary(node *ast.BinaryNode) {
	switch node.Operator {
	case "+":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpAdd)

	case "-":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpSub)

	case "*":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpMul)

	case "/":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpDiv)

	case "%":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpMod)

	case "^":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpExp)

	case "==":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		// TODO: check expr
		compiler.emit(code.OpEqual)

	case "!=":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpNotEqual)

	case ">":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpGreaterThan)

	case "<":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpLessThan)

	case ">=":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpGreaterOrEqual)

	case "<=":
		compiler.compile(node.Left)
		compiler.compile(node.Right)
		compiler.emit(code.OpLessOrEqual)

	case "or":
		compiler.compile(node.Left)
		compiler.emit(code.OpJumpIfTrue)
		compiler.emit(code.OpPop)
		compiler.compile(node.Right)

	case "and":
		compiler.compile(node.Left)
		compiler.emit(code.OpJumpIfFalse)
		compiler.emit(code.OpPop)
		compiler.compile(node.Right)
	}
	// compiler.emit(code.OpPop)
}

func (compiler *Compiler) patchJump(placeholder int) {
	//compiler.arguments[placeholder-1] = len(compiler.instructions) - placeholder
}

func (compiler *Compiler) NodeFunction(node *ast.FunctionNode) {
	// TODO: implement
}

func (compiler *Compiler) NodeArray(node *ast.ArrayNode) {
	for _, node := range node.Nodes {
		compiler.compile(node)
	}
	compiler.emit(code.OpArray, len(node.Nodes))
}

func (compiler *Compiler) addInstruction(ins code.Instructions) int {
	compiler.instructions = append(compiler.instructions, ins)
	posNewInstruction := len(compiler.instructions)
	return posNewInstruction
}

func (compiler *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := compiler.addInstruction(ins)
	return pos
}

func (compiler *Compiler) addConstant(constant interface{}) int {
	compiler.constants = append(compiler.constants, constant)
	return len(compiler.constants) - 1
}

package ast

type NodeType int

type Node interface{}

type NumberNode struct {
	NodeType
	Value   string
	IsInt   bool
	Int64   int64
	IsFloat bool
	Float64 float64
}

type IdentifierNode struct {
	NodeType
	Value string
}

type StringNode struct {
	NodeType
	Value string
}

type BoolNode struct {
	NodeType
	Value bool
}

type NilNode struct {
	NodeType
}

type UnaryNode struct {
	Operator string
	Node     Node
}

type BinaryNode struct {
	Operator string
	Left     Node
	Right    Node
}

const (
	NodeText NodeType = iota
	NodeNumber
	NodeIdentifier
	NodeString
	NodeBool
	NodeNil
)

package ast

type NodeType int

type Node interface {
	Type() NodeType
}

func (t NodeType) Type() NodeType {
	return t
}

type NumberNode struct {
	NodeType
	Value   string
	IsInt   bool
	Int64   int64
	IsFloat bool
	Float64 float64
}

func (node *NumberNode) Type() NodeType {
	return NodeNumber
}

type IdentifierNode struct {
	NodeType
	Value string
}

func (node *IdentifierNode) Type() NodeType {
	return NodeIdentifier
}

type StringNode struct {
	NodeType
	Value string
}

func (node *StringNode) Type() NodeType {
	return NodeString
}

type BoolNode struct {
	NodeType
	Value bool
}

func (node *BoolNode) Type() NodeType {
	return NodeBool
}

type NilNode struct {
	NodeType
}

func (node *NilNode) Type() NodeType {
	return NodeNil
}

type UnaryNode struct {
	NodeType
	Operator string
	Node     Node
}

func (node *UnaryNode) Type() NodeType {
	return NodeUnary
}

type BinaryNode struct {
	NodeType
	Operator string
	Left     Node
	Right    Node
}

func (node *BinaryNode) Type() NodeType {
	return NodeBinary
}

type FunctionNode struct {
	NodeType
	Function  Node
	Arguments []Node
}

type ArrayNode struct {
	NodeType
	Nodes []Node
}

type MemberNode struct {
	NodeType
	Node     Node
	Property Node
}

const (
	NodeNumber NodeType = iota
	NodeIdentifier
	NodeString
	NodeBool
	NodeNil
	NodeUnary
	NodeBinary
	NodeFunction
	NodeArray
	NodeMember
)

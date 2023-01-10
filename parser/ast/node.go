package ast

type NodeType int

type Node interface{}

type NumberNode struct {
	// NodeType
	Value string
}

const (
	NodeText NodeType = iota
	NodeNumber
)

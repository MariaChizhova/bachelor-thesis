package compiler

import (
	"bachelor-thesis/vm/code"
)

// Program TODO: move to vm
type Program struct {
	Instructions code.Instructions
	Constants    []interface{}
}

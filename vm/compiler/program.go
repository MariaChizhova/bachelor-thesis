package compiler

import (
	"bachelor-thesis/code"
)

// Program TODO: move to vm
type Program struct {
	Instructions code.Instructions
	Constants    []interface{}
}

package main

import (
	"bachelor-thesis/evaluator"
	"bachelor-thesis/parser"
	"bachelor-thesis/vm"
	"bachelor-thesis/vm/compiler"
	"bachelor-thesis/vm2"
	"bachelor-thesis/vm3"
	"fmt"
)

type vmType int64

const (
	treeTraversal vmType = iota
	singleStack
	multipleStacks
	reflectBased
	register
)

func Eval(input string, vmType vmType, env interface{}) (interface{}, error) {
	tree := parser.Parse(input)
	if vmType == treeTraversal {
		evaluated, err := evaluator.Eval(tree)
		if err != nil {
			return nil, err
		}
		return evaluated, nil
	} else {
		program, err := compiler.Compile(tree)
		if err != nil {
			return nil, err
		}
		if vmType == singleStack {
			vm := vm.New(program.Instructions, program.Constants)
			err = vm.Run()
			if err != nil {
				return nil, err
			}
			return vm.StackTop(), nil
		} else if vmType == multipleStacks {
			vm := vm2.New(program.Instructions, program.Constants)
			err = vm.Run()
			if err != nil {
				return nil, err
			}
			return vm.StackTop(), nil
		} else if vmType == reflectBased {
			vm := vm3.New(program.Instructions, program.Constants)
			err = vm.Run()
			if err != nil {
				return nil, err
			}
			return vm.StackTop(), nil
		} else if vmType == register {
			return nil, fmt.Errorf("register based machine is not implemented now")
		}
	}
	return nil, fmt.Errorf("undefined type of the virtual machine")
}

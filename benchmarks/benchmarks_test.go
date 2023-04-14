package benchmarks

import (
	"bachelor-thesis/evaluator"
	"bachelor-thesis/parser"
	"bachelor-thesis/vm"
	"bachelor-thesis/vm/compiler"
	"bachelor-thesis/vm2"
	"bachelor-thesis/vm3"
	"strconv"
	"testing"
)

func getSum(n int) string {
	out := "1 "
	for i := 2; i <= n; i++ {
		out += " + " + strconv.Itoa(i)
	}
	return out
}

func Benchmark_treeTraversal(b *testing.B) {
	tree := parser.Parse(getSum(100))
	var out interface{}
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = evaluator.Eval(tree)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 5050 {
		b.Fail()
	}
}

func Benchmark_singleStack(b *testing.B) {
	tree := parser.Parse(getSum(100))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm.New(program.Instructions, program.Constants)
		err = vm.Run()
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 5050 {
		b.Fail()
	}
}

func Benchmark_multipleStacks(b *testing.B) {
	tree := parser.Parse(getSum(100))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm2.New(program.Instructions, program.Constants)
		err = vm.Run()
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 5050 {
		b.Fail()
	}
}

func Benchmark_reflectBased(b *testing.B) {
	tree := parser.Parse(getSum(100))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm3.New(program.Instructions, program.Constants)
		err = vm.Run()
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 5050 {
		b.Fail()
	}
}

func Benchmark_singleStackStrings(b *testing.B) {
	tree := parser.Parse(`"a" + "b" + "c" + "d" + "e" + "d" + "e"`)
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm.New(program.Instructions, program.Constants)
		err = vm.Run()
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(string) != "abcdede" {
		b.Fail()
	}
}

func Benchmark_multipleStacksStrings(b *testing.B) {
	tree := parser.Parse(`"a" + "b" + "c" + "d" + "e" + "d" + "e"`)
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm2.New(program.Instructions, program.Constants)
		err = vm.Run()
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(string) != "abcdede" {
		b.Fail()
	}
}

package benchmarks

import (
	"bachelor-thesis/evaluator"
	"bachelor-thesis/parser"
	"bachelor-thesis/vm"
	"bachelor-thesis/vm/compiler"
	"bachelor-thesis/vm2"
	"bachelor-thesis/vm3"
	"bachelor-thesis/vm4"
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
	tree := parser.Parse(getSum(5))
	var out interface{}
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = evaluator.Eval(tree, nil)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 15 {
		b.Fail()
	}
}

func Benchmark_singleStack(b *testing.B) {
	tree := parser.Parse(getSum(5))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm.New(program.Instructions, program.Constants)
		err = vm.Run(nil)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 15 {
		b.Fail()
	}
}

func Benchmark_multipleStacks(b *testing.B) {
	tree := parser.Parse(getSum(5))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm2.New(program.Instructions, program.Constants)
		err = vm.Run(nil)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 15 {
		b.Fail()
	}
}

func Benchmark_reflectBased(b *testing.B) {
	tree := parser.Parse(getSum(5))
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm3.New(program.Instructions, program.Constants)
		err = vm.Run(nil)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 15 {
		b.Fail()
	}
}

func Benchmark_registerBased(b *testing.B) {
	program := []int64{
		vm4.OpConstant, vm4.R0, 1,
		vm4.OpConstant, vm4.R1, 2,
		vm4.OpAdd, vm4.R0, vm4.R1,
		vm4.OpConstant, vm4.R1, 3,
		vm4.OpAdd, vm4.R0, vm4.R1,
		vm4.OpConstant, vm4.R1, 4,
		vm4.OpAdd, vm4.R0, vm4.R1,
		vm4.OpConstant, vm4.R1, 5,
		vm4.OpAdd, vm4.R0, vm4.R1,
		vm4.OpPrint, vm4.R0,
		vm4.OpHalt}
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm4.New(program)
		vm.Run()
		out = vm.GetResult()
	}
	b.StopTimer()
	if out.(int64) != 15 {
		b.Fail()
	}
}

// Strings
func Benchmark_singleStackStrings(b *testing.B) {
	tree := parser.Parse(`"a" + "b" + "c" + "d" + "e" + "d" + "e"`)
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm.New(program.Instructions, program.Constants)
		err = vm.Run(nil)
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
		err = vm.Run(nil)
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

// Function calls
func Benchmark_treeTraversalCalls(b *testing.B) {
	env := map[string]interface{}{"add": func(a, b int64) int64 { return a + b }}
	tree := parser.Parse("add(1, 2)")
	var out interface{}
	var err error
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = evaluator.Eval(tree, env)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 3 {
		b.Fail()
	}
}

func Benchmark_singleStackCalls(b *testing.B) {
	env := map[string]interface{}{"add": func(a, b int64) int64 { return a + b }}
	tree := parser.Parse("add(1, 2)")
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm.New(program.Instructions, program.Constants)
		err = vm.Run(env)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 3 {
		b.Fail()
	}
}

func Benchmark_multipleStacksCalls(b *testing.B) {
	env := map[string]interface{}{"add": func(a, b int64) int64 { return a + b }}
	tree := parser.Parse("add(1, 2)")
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm2.New(program.Instructions, program.Constants)
		err = vm.Run(env)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 3 {
		b.Fail()
	}
}

func Benchmark_reflectBasedCalls(b *testing.B) {
	env := map[string]interface{}{"add": func(a, b int64) int64 { return a + b }}
	tree := parser.Parse("add(1, 2)")
	program, err := compiler.Compile(tree)
	var out interface{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := vm3.New(program.Instructions, program.Constants)
		err = vm.Run(env)
		out = vm.StackTop()
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 3 {
		b.Fail()
	}
}

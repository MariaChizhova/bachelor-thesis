package benchmarks

import (
	"bachelor-thesis/evaluator"
	"bachelor-thesis/parser"
	"bachelor-thesis/vm"
	"bachelor-thesis/vm/compiler"
	"bachelor-thesis/vm2"
	"bachelor-thesis/vm3"
	"bachelor-thesis/vm4"
	"bachelor-thesis/vm5"
	"fmt"
	"github.com/antonmedv/expr"
	"math"
	"testing"
)

func Benchmark_treeTraversal(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getSum(i))
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
			if out.(int64) != int64(i*(i+1)/2) {
				b.Fail()
			}
		})
	}
}

func Benchmark_singleStack(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getSum(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm.New(program.Instructions, program.Constants)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(int64) != int64(i*(i+1)/2) {
				b.Fail()
			}
		})
	}
}

func Benchmark_multipleStacks(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getSum(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm2.New(program.Instructions, program.Constants)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(int64) != int64(i*(i+1)/2) {
				b.Fail()
			}
		})
	}
}

func Benchmark_reflectBased(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getSum(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm3.New(program.Instructions, program.Constants)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(int64) != int64(i*(i+1)/2) {
				b.Fail()
			}
		})
	}
}

func Benchmark_registerBased(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateSumBytecode(i)
			var out interface{}
			vm := vm4.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run()
			}
			b.StopTimer()
			out = vm.GetResult()
			if out.(int64) != int64(i*(i+1)/2) {
				b.Fail()
			}
		})
	}
}

func Benchmark_registerBasedInterfaces(b *testing.B) {
	for i := 1; i <= 76; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateSumBytecode2(i)
			vm := vm5.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run()
			}
			b.StopTimer()
			out := vm.Registers[3].(int)
			if out != i*(i+1)/2 {
				b.Fail()
			}
		})
	}
}

// Strings
func Benchmark_treeTraversalStrings(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(concatenateStrings(i))
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
			if out.(string) != concatenateStringsResult(i) {
				b.Fail()
			}
		})
	}
}

func Benchmark_singleStackStrings(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(concatenateStrings(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm.New(program.Instructions, program.Constants)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(string) != concatenateStringsResult(i) {
				b.Fail()
			}
		})
	}
}

func Benchmark_multipleStacksStrings(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(concatenateStrings(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm2.New(program.Instructions, program.Constants)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(string) != concatenateStringsResult(i) {
				b.Fail()
			}
		})
	}
}

func Benchmark_reflectBasedStrings(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(concatenateStrings(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm3.New(program.Instructions, program.Constants)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()

			if err != nil {
				b.Fatal(err)
			}
			if out.(string) != concatenateStringsResult(i) {
				b.Fail()
			}
		})
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
	vm := vm.New(program.Instructions, program.Constants)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = vm.Run(env)
	}
	b.StopTimer()
	out = vm.StackTop()

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
	vm := vm2.New(program.Instructions, program.Constants)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = vm.Run(env)
	}
	b.StopTimer()
	out = vm.StackTop()

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
	vm := vm3.New(program.Instructions, program.Constants)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = vm.Run(env)
	}
	b.StopTimer()
	out = vm.StackTop()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != 3 {
		b.Fail()
	}
}

// expression
func Benchmark_stackBasedExpression(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getExpression(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm3.New(program.Instructions, program.Constants)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
			}
			b.StopTimer()
			out = vm.StackTop()
			if err != nil {
				b.Fatal(err)
			}
			if i%2 != 0 {
				if out.(int64) != int64((i+1)/2) {
					b.Fail()
				}
			} else {
				if out.(int64) != int64(-i/2) {
					b.Fail()
				}
			}
		})
	}
}

func Benchmark_registerBasedExpression(b *testing.B) {
	for i := 1; i <= 100; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateExpressionBytecode2(i)
			vm := vm5.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run()
			}
			b.StopTimer()
			out := vm.Registers[3].(int)
			if i%2 != 0 {
				if out != (i+1)/2 {
					b.Fail()
				}
			} else {
				if out != -i/2 {
					b.Fail()
				}
			}
		})
	}
}

// different number of functions with the same arguments
var env = map[string]interface{}{
	"foo1":  func(a, b int64) int64 { return a + b },
	"foo2":  func(a, b int64) int64 { return a - b },
	"foo3":  func(a, b int64) int64 { return a * b },
	"foo4":  func(a, b int64) int64 { return a / b },
	"foo5":  func(a, b int64) int64 { return a % b },
	"foo6":  func(a, b int64) int64 { return int64(math.Pow(float64(a), float64(b))) },
	"foo7":  func(a, b int64) int64 { return 2*a + b },
	"foo8":  func(a, b int64) int64 { return a + 2*b },
	"foo9":  func(a, b int64) int64 { return 2*a + 2*b },
	"foo10": func(a, b int64) int64 { return 2*a - 2*b },
}
var code = "foo1(1, 2) + foo2(1, 2) + foo3(1, 2) + foo4(4, 2) + foo5(5, 2) + foo6(2, 2) + foo7(1, 2) + foo8(1, 2) + foo9(1, 2) + foo10(1, 1)"
var result = 26

func Benchmark_1(b *testing.B) {
	tree := parser.Parse(code)
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
	if out.(int64) != int64(result) {
		print(out.(int64))
		b.Fail()
	}
}

func Benchmark_2(b *testing.B) {
	tree := parser.Parse(code)
	program, err := compiler.Compile(tree)
	var out interface{}
	vm := vm.New(program.Instructions, program.Constants)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = vm.Run(env)
	}
	b.StopTimer()
	out = vm.StackTop()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != int64(result) {
		b.Fail()
	}
}

func Benchmark_3(b *testing.B) {
	tree := parser.Parse(code)
	program, err := compiler.Compile(tree)
	var out interface{}
	vm := vm3.New(program.Instructions, program.Constants)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = vm.Run(env)
	}
	b.StopTimer()
	out = vm.StackTop()

	if err != nil {
		b.Fatal(err)
	}
	if out.(int64) != int64(result) {
		b.Fail()
	}
}

// combination of booleans and strings like these:
// ("a" > "b") and ("a" == "c") or ("x" <= "xy"), where numArgs = number of Brackets
func Benchmark_booleansStrings(b *testing.B) {
	for i := 1; i <= 50; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			code = generateString(i)
			tree := parser.Parse(code)
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
			res, _ := expr.Eval(code, nil)
			if out.(bool) != res {
				b.Fail()
			}
		})
	}
}

package benchmarks

import (
	"bachelor-thesis/evaluator"
	"bachelor-thesis/parser"
	"bachelor-thesis/vm"
	"bachelor-thesis/vm/compiler"
	"bachelor-thesis/vm2"
	"bachelor-thesis/vm3"
	"bachelor-thesis/vm5"
	"bachelor-thesis/vm7"
	"fmt"
	"github.com/antonmedv/expr"
	"math"
	"testing"
)

func Benchmark_treeTraversal(b *testing.B) {
	for i := 200; i <= 200; i++ {
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
	for i := 200; i <= 200; i++ {
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
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getSum(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm7.New(program.Instructions, program.Constants)

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
	for i := 200; i <= 200; i++ {
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

func Benchmark_registerBasedSum(b *testing.B) {
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateSumBytecode(i)
			vm := vm5.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run(nil)
			}
			b.StopTimer()
			out := vm.Registers[3]
			if out != i*(i+1)/2 {
				b.Fail()
			}
		})
	}
}

// Strings
func Benchmark_treeTraversalStrings(b *testing.B) {
	for i := 200; i <= 200; i++ {
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
	for i := 200; i <= 200; i++ {
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
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(concatenateStrings(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm7.New(program.Instructions, program.Constants)
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
	for i := 200; i <= 200; i++ {
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

func Benchmark_registerBasedStrings(b *testing.B) {
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateBytecodeStrings(i)
			var out interface{}
			vm := vm5.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run(nil)
			}
			b.StopTimer()
			out = vm.Registers[3]
			result := concatenateStringsResult(i)
			if out != result {
				fmt.Println(out)
				fmt.Println(result)
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

func Benchmark_registerBasedCalls(b *testing.B) {
	env := map[string]interface{}{"add": func(a, b, c, d, e, f, g, k, l, m int) int { return a + b + c + d + e + f + g + k + l + m }}
	program := vm5.Program{
		Instructions: []byte{
			byte(vm5.OpStoreInt), 01, 0,
			byte(vm5.OpStoreInt), 02, 1,
			byte(vm5.OpStoreInt), 03, 2,
			byte(vm5.OpStoreInt), 04, 3,
			byte(vm5.OpStoreInt), 05, 4,
			byte(vm5.OpStoreInt), 06, 5,
			byte(vm5.OpStoreInt), 07, 6,
			byte(vm5.OpStoreInt), 8, 7,
			byte(vm5.OpStoreInt), 9, 8,
			byte(vm5.OpStoreInt), 10, 9,
			byte(vm5.OpCall), 03, 10, 10, 01, 02, 03, 04, 05, 06, 07, 8, 9, 10,
		},
		Constants: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, "add"},
	}
	var out interface{}
	vm := vm5.New(program)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm.Run(env)
	}
	b.StopTimer()
	out = vm.Registers[3]
	if out != 55 {
		b.Fail()
	}
}

// expression
func Benchmark_stackBasedExpression(b *testing.B) {
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			tree := parser.Parse(getExpression(i))
			program, err := compiler.Compile(tree)
			var out interface{}
			//var err error
			vm := vm7.New(program.Instructions, program.Constants)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(nil)
				//out, err = evaluator.Eval(tree, nil)
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
	for i := 200; i <= 200; i++ {
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			program := generateExpressionBytecode(i)
			vm := vm5.New(program)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				vm.Run(nil)
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

var env1 = map[string]interface{}{
	"foo1":  func(a, b int) int { return a + b },
	"foo2":  func(a, b int) int { return a - b },
	"foo3":  func(a, b int) int { return a * b },
	"foo4":  func(a, b int) int { return a / b },
	"foo5":  func(a, b int) int { return a % b },
	"foo6":  func(a, b int) int { return int(math.Pow(float64(a), float64(b))) },
	"foo7":  func(a, b int) int { return 2*a + b },
	"foo8":  func(a, b int) int { return a + 2*b },
	"foo9":  func(a, b int) int { return 2*a + 2*b },
	"foo10": func(a, b int) int { return 2*a - 2*b },
}

func Benchmark_registerBasedCalls2(b *testing.B) {
	program := vm5.Program{
		Instructions: []byte{
			byte(vm5.OpStoreInt), 01, 0,
			byte(vm5.OpStoreInt), 02, 1,
			byte(vm5.OpCall), 03, 2, 2, 01, 02,

			byte(vm5.OpStoreInt), 01, 3,
			byte(vm5.OpStoreInt), 02, 4,
			byte(vm5.OpCall), 04, 5, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 6,
			byte(vm5.OpStoreInt), 02, 7,
			byte(vm5.OpCall), 04, 8, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 9,
			byte(vm5.OpStoreInt), 02, 10,
			byte(vm5.OpCall), 04, 11, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 12,
			byte(vm5.OpStoreInt), 02, 13,
			byte(vm5.OpCall), 04, 14, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 15,
			byte(vm5.OpStoreInt), 02, 16,
			byte(vm5.OpCall), 04, 17, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 18,
			byte(vm5.OpStoreInt), 02, 19,
			byte(vm5.OpCall), 04, 20, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 21,
			byte(vm5.OpStoreInt), 02, 22,
			byte(vm5.OpCall), 04, 23, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 24,
			byte(vm5.OpStoreInt), 02, 25,
			byte(vm5.OpCall), 04, 26, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,

			byte(vm5.OpStoreInt), 01, 27,
			byte(vm5.OpStoreInt), 02, 28,
			byte(vm5.OpCall), 04, 29, 2, 01, 02,
			byte(vm5.OpAdd), 03, 03, 04,
		},
		Constants: []interface{}{1, 2, "foo1", 1, 2, "foo2", 1, 2, "foo3", 4, 2, "foo4", 5, 2, "foo5", 2, 2, "foo6",
			1, 2, "foo7", 1, 2, "foo8", 1, 2, "foo9", 1, 2, "foo10"}}
	vm := vm5.New(program)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm.Run(env1)
	}
	b.StopTimer()
	out := vm.Registers[3]
	result := (1 + 2) + (1 - 2) + (1 * 2) + (4 / 2) + (5 % 2) + int(math.Pow(float64(2), float64(2))) + (2*1 + 2) + (1 + 2*2) + (2*1 + 2*2) + (2*1 - 2*2)
	if out != result {
		b.Fail()
	}
}

// combination of booleans and strings like these:
// ("a" > "b") and ("a" == "c") or ("x" <= "xy"), where numArgs = number of Brackets
func Benchmark_booleansStrings11(b *testing.B) {
	a := "(\"oe\" == \"e\") and (\"a\" != \"kmad\") and (\"iq\" == \"fpbx\") and (\"z\" != \"tgf\") and (\"a\" >= \"yc\") and (\"bh\" != \"gu\") and (\"zc\" <= \"a\") and (\"i\" >= \"ot\") and (\"utujv\" <= \"ubq\") and (\"pca\" < \"vw\") and (\"bcea\" >= \"uxuq\") and (\"ddf\" <= \"tzp\") and (\"ky\" != \"yegm\") and (\"ywyej\" <= \"bw\") and (\"be\" < \"qban\") and (\"hbgr\" != \"sqe\") and (\"tmc\" > \"ifux\") and (\"didgv\" < \"l\") and (\"tfa\" < \"ij\") and (\"dyvrs\" < \"yijkw\") and (\"sdv\" >= \"jriux\") and (\"jd\" <= \"yxe\") and (\"es\" <= \"xb\") and (\"czrug\" <= \"w\") and (\"khydq\" != \"gar\") and (\"ot\" < \"fl\") and (\"fqmaf\" != \"thr\") and (\"rdg\" != \"e\") and (\"yahb\" != \"khdsk\") and (\"axgi\" == \"bx\") and (\"thln\" > \"hfh\") and (\"jny\" == \"bjomw\") and (\"lc\" == \"tce\") and (\"zok\" >= \"u\") and (\"xmfl\" > \"schcy\") and (\"kt\" > \"zvni\") and (\"svtj\" >= \"q\") and (\"x\" < \"lae\") and (\"zuetf\" > \"tpkc\") and (\"nowy\" == \"gaa\") and (\"zk\" <= \"bbjz\") and (\"zexc\" <= \"v\") and (\"bm\" <= \"jhmjs\") and (\"zl\" < \"cbhp\") and (\"ixrsj\" >= \"cy\") and (\"iet\" != \"b\") and (\"pprq\" == \"tqqls\") and (\"n\" == \"khwdr\") and (\"co\" >= \"apmt\") and (\"pfgoq\" >= \"rvfm\") and (\"qrlv\" == \"zkdr\") and (\"gpoqy\" == \"qbfig\") and (\"i\" <= \"fplix\") and (\"ckcu\" >= \"r\") and (\"mequy\" != \"nu\") and (\"z\" >= \"p\") and (\"qiqge\" < \"jkwcy\") and (\"s\" <= \"iiedc\") and (\"fkcz\" <= \"b\") and (\"bbd\" <= \"bvzg\") and (\"wq\" > \"jdua\") and (\"z\" == \"pqqz\") and (\"pw\" != \"djbp\") and (\"sriu\" != \"dctvv\") and (\"ak\" == \"qmhe\") and (\"nwipo\" < \"jwe\") and (\"dr\" >= \"ec\") and (\"jt\" >= \"gus\") and (\"lo\" != \"maqe\") and (\"ee\" != \"tiyq\") and (\"qqdd\" <= \"osih\") and (\"feraf\" < \"e\") and (\"jh\" >= \"c\") and (\"o\" > \"qh\") and (\"qzrkj\" == \"pyiy\") and (\"fdv\" == \"luu\") and (\"emyp\" >= \"fj\") and (\"goij\" > \"usyz\") and (\"hx\" == \"da\") and (\"zteu\" < \"z\") and (\"iicb\" < \"e\") and (\"zely\" == \"xsal\") and (\"uw\" != \"zymi\") and (\"rrluj\" != \"k\") and (\"o\" >= \"lg\") and (\"tzjut\" == \"p\") and (\"k\" <= \"pb\") and (\"hnng\" >= \"iwu\") and (\"ucpwg\" > \"wzx\") and (\"ttxnj\" < \"fl\") and (\"juu\" >= \"bx\") and (\"yf\" > \"i\") and (\"bnpt\" <= \"rlm\") and (\"ciz\" < \"ggg\") and (\"jlv\" != \"n\") and (\"hpup\" < \"g\") and (\"dxe\" >= \"awdg\") and (\"acbk\" >= \"nzm\") and (\"uki\" >= \"fd\") and (\"ti\" >= \"aryb\") and (\"ga\" > \"r\") and (\"d\" >= \"qawqn\") and (\"xup\" != \"pum\") and (\"h\" != \"azjl\") and (\"kw\" != \"uuxxw\") and (\"syfl\" != \"oisp\") and (\"ath\" < \"s\") and (\"s\" >= \"sgwnw\") and (\"esq\" == \"tfz\") and (\"w\" <= \"cla\") and (\"k\" < \"nja\") and (\"e\" != \"wscvd\") and (\"wln\" < \"it\") and (\"qefl\" < \"cqdbu\") and (\"on\" > \"tabg\") and (\"xpde\" == \"a\") and (\"ol\" != \"ie\") and (\"bcweo\" >= \"detbm\") and (\"zf\" != \"wsj\") and (\"r\" >= \"qtn\") and (\"wjyt\" > \"ev\") and (\"g\" < \"uirg\") and (\"v\" > \"z\") and (\"xfj\" >= \"dpn\") and (\"fyq\" > \"pcbve\") and (\"wh\" >= \"wlh\") and (\"seh\" == \"jflvz\") and (\"t\" >= \"pqjz\") and (\"srcqz\" < \"ysn\") and (\"ytxoa\" <= \"rgqkf\") and (\"llgf\" == \"wum\") and (\"rcggd\" == \"do\") and (\"mtemz\" != \"ib\") and (\"h\" >= \"oym\") and (\"duqq\" == \"dv\") and (\"pd\" < \"irasf\") and (\"yj\" >= \"m\") and (\"fg\" >= \"sohlk\") and (\"dwjlz\" >= \"r\") and (\"hunl\" < \"rwe\") and (\"c\" >= \"pe\") and (\"dqawy\" == \"b\") and (\"llb\" != \"wymue\") and (\"li\" >= \"q\") and (\"dsxeg\" < \"xasgn\") and (\"f\" > \"x\") and (\"mj\" < \"d\") and (\"dju\" >= \"u\") and (\"qem\" == \"d\") and (\"pdpjx\" >= \"fxr\") and (\"lqvmd\" > \"sc\") and (\"mgytq\" > \"aggif\") and (\"tbx\" < \"jedpa\") and (\"ru\" != \"zktci\") and (\"nt\" <= \"i\") and (\"dxjw\" > \"rlq\") and (\"riwe\" > \"yxg\") and (\"vlrj\" != \"jdor\") and (\"ps\" != \"bavfw\") and (\"lal\" == \"jse\") and (\"tdl\" <= \"vlie\") and (\"fo\" >= \"lag\") and (\"t\" < \"ql\") and (\"nq\" >= \"mcll\") and (\"gfv\" != \"jclb\") and (\"ctbb\" > \"u\") and (\"vzk\" >= \"ky\") and (\"wtxy\" >= \"tgsly\") and (\"x\" >= \"j\") and (\"z\" == \"b\") and (\"xav\" <= \"lhmq\") and (\"t\" <= \"yqom\") and (\"wkmzt\" > \"dt\") and (\"eru\" > \"r\") and (\"nb\" >= \"vkxyo\") and (\"avsvz\" == \"z\") and (\"weaj\" > \"iazr\") and (\"cdx\" < \"vqrd\") and (\"svlm\" > \"zcvsd\") and (\"n\" == \"sa\") and (\"vt\" > \"qg\") and (\"av\" > \"o\") and (\"hulcd\" != \"k\") and (\"fasnc\" == \"iw\") and (\"t\" < \"vk\") and (\"l\" == \"neh\") and (\"iplv\" != \"fmo\") and (\"g\" < \"lewi\") and (\"xt\" >= \"o\") and (\"uivwz\" == \"fagls\") and (\"oyad\" == \"gri\") and (\"d\" > \"ec\") and (\"rpztf\" >= \"bqtfb\") and (\"hm\" >= \"yfrni\") and (\"pjz\" >= \"gnnd\") and (\"clw\" <= \"zkpd\") and (\"casr\" >= \"pvw\") and (\"vmcvd\" <= \"rvabv\") and (\"zj\" < \"wyfr\") and (\"ki\" <= \"l\")"
	for i := 0; i < 1; i++ {
		code = a
		b.Run(fmt.Sprintf("input-%d", i), func(b *testing.B) {
			//code = generateString(i)
			tree := parser.Parse(code)
			//var out interface{}
			//var err error
			//b.ResetTimer()
			//for n := 0; n < b.N; n++ {
			//	out, err = evaluator.Eval(tree, nil)
			//}

			program, err := compiler.Compile(tree)
			var out interface{}
			vm := vm7.New(program.Instructions, program.Constants)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				err = vm.Run(env)
			}
			b.StopTimer()
			out = vm.StackTop()
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

func Benchmark_booleansStringsRegister(b *testing.B) {
	a := "(\"oe\" == \"e\") and (\"a\" != \"kmad\") and (\"iq\" == \"fpbx\") and (\"z\" != \"tgf\") and (\"a\" >= \"yc\") and (\"bh\" != \"gu\") and (\"zc\" <= \"a\") and (\"i\" >= \"ot\") and (\"utujv\" <= \"ubq\") and (\"pca\" < \"vw\") and (\"bcea\" >= \"uxuq\") and (\"ddf\" <= \"tzp\") and (\"ky\" != \"yegm\") and (\"ywyej\" <= \"bw\") and (\"be\" < \"qban\") and (\"hbgr\" != \"sqe\") and (\"tmc\" > \"ifux\") and (\"didgv\" < \"l\") and (\"tfa\" < \"ij\") and (\"dyvrs\" < \"yijkw\") and (\"sdv\" >= \"jriux\") and (\"jd\" <= \"yxe\") and (\"es\" <= \"xb\") and (\"czrug\" <= \"w\") and (\"khydq\" != \"gar\") and (\"ot\" < \"fl\") and (\"fqmaf\" != \"thr\") and (\"rdg\" != \"e\") and (\"yahb\" != \"khdsk\") and (\"axgi\" == \"bx\") and (\"thln\" > \"hfh\") and (\"jny\" == \"bjomw\") and (\"lc\" == \"tce\") and (\"zok\" >= \"u\") and (\"xmfl\" > \"schcy\") and (\"kt\" > \"zvni\") and (\"svtj\" >= \"q\") and (\"x\" < \"lae\") and (\"zuetf\" > \"tpkc\") and (\"nowy\" == \"gaa\") and (\"zk\" <= \"bbjz\") and (\"zexc\" <= \"v\") and (\"bm\" <= \"jhmjs\") and (\"zl\" < \"cbhp\") and (\"ixrsj\" >= \"cy\") and (\"iet\" != \"b\") and (\"pprq\" == \"tqqls\") and (\"n\" == \"khwdr\") and (\"co\" >= \"apmt\") and (\"pfgoq\" >= \"rvfm\") and (\"qrlv\" == \"zkdr\") and (\"gpoqy\" == \"qbfig\") and (\"i\" <= \"fplix\") and (\"ckcu\" >= \"r\") and (\"mequy\" != \"nu\") and (\"z\" >= \"p\") and (\"qiqge\" < \"jkwcy\") and (\"s\" <= \"iiedc\") and (\"fkcz\" <= \"b\") and (\"bbd\" <= \"bvzg\") and (\"wq\" > \"jdua\") and (\"z\" == \"pqqz\") and (\"pw\" != \"djbp\") and (\"sriu\" != \"dctvv\") and (\"ak\" == \"qmhe\") and (\"nwipo\" < \"jwe\") and (\"dr\" >= \"ec\") and (\"jt\" >= \"gus\") and (\"lo\" != \"maqe\") and (\"ee\" != \"tiyq\") and (\"qqdd\" <= \"osih\") and (\"feraf\" < \"e\") and (\"jh\" >= \"c\") and (\"o\" > \"qh\") and (\"qzrkj\" == \"pyiy\") and (\"fdv\" == \"luu\") and (\"emyp\" >= \"fj\") and (\"goij\" > \"usyz\") and (\"hx\" == \"da\") and (\"zteu\" < \"z\") and (\"iicb\" < \"e\") and (\"zely\" == \"xsal\") and (\"uw\" != \"zymi\") and (\"rrluj\" != \"k\") and (\"o\" >= \"lg\") and (\"tzjut\" == \"p\") and (\"k\" <= \"pb\") and (\"hnng\" >= \"iwu\") and (\"ucpwg\" > \"wzx\") and (\"ttxnj\" < \"fl\") and (\"juu\" >= \"bx\") and (\"yf\" > \"i\") and (\"bnpt\" <= \"rlm\") and (\"ciz\" < \"ggg\") and (\"jlv\" != \"n\") and (\"hpup\" < \"g\") and (\"dxe\" >= \"awdg\") and (\"acbk\" >= \"nzm\") and (\"uki\" >= \"fd\") and (\"ti\" >= \"aryb\") and (\"ga\" > \"r\") and (\"d\" >= \"qawqn\") and (\"xup\" != \"pum\") and (\"h\" != \"azjl\") and (\"kw\" != \"uuxxw\") and (\"syfl\" != \"oisp\") and (\"ath\" < \"s\") and (\"s\" >= \"sgwnw\") and (\"esq\" == \"tfz\") and (\"w\" <= \"cla\") and (\"k\" < \"nja\") and (\"e\" != \"wscvd\") and (\"wln\" < \"it\") and (\"qefl\" < \"cqdbu\") and (\"on\" > \"tabg\") and (\"xpde\" == \"a\") and (\"ol\" != \"ie\") and (\"bcweo\" >= \"detbm\") and (\"zf\" != \"wsj\") and (\"r\" >= \"qtn\") and (\"wjyt\" > \"ev\") and (\"g\" < \"uirg\") and (\"v\" > \"z\") and (\"xfj\" >= \"dpn\") and (\"fyq\" > \"pcbve\") and (\"wh\" >= \"wlh\") and (\"seh\" == \"jflvz\") and (\"t\" >= \"pqjz\") and (\"srcqz\" < \"ysn\") and (\"ytxoa\" <= \"rgqkf\") and (\"llgf\" == \"wum\") and (\"rcggd\" == \"do\") and (\"mtemz\" != \"ib\") and (\"h\" >= \"oym\") and (\"duqq\" == \"dv\") and (\"pd\" < \"irasf\") and (\"yj\" >= \"m\") and (\"fg\" >= \"sohlk\") and (\"dwjlz\" >= \"r\") and (\"hunl\" < \"rwe\") and (\"c\" >= \"pe\") and (\"dqawy\" == \"b\") and (\"llb\" != \"wymue\") and (\"li\" >= \"q\") and (\"dsxeg\" < \"xasgn\") and (\"f\" > \"x\") and (\"mj\" < \"d\") and (\"dju\" >= \"u\") and (\"qem\" == \"d\") and (\"pdpjx\" >= \"fxr\") and (\"lqvmd\" > \"sc\") and (\"mgytq\" > \"aggif\") and (\"tbx\" < \"jedpa\") and (\"ru\" != \"zktci\") and (\"nt\" <= \"i\") and (\"dxjw\" > \"rlq\") and (\"riwe\" > \"yxg\") and (\"vlrj\" != \"jdor\") and (\"ps\" != \"bavfw\") and (\"lal\" == \"jse\") and (\"tdl\" <= \"vlie\") and (\"fo\" >= \"lag\") and (\"t\" < \"ql\") and (\"nq\" >= \"mcll\") and (\"gfv\" != \"jclb\") and (\"ctbb\" > \"u\") and (\"vzk\" >= \"ky\") and (\"wtxy\" >= \"tgsly\") and (\"x\" >= \"j\") and (\"z\" == \"b\") and (\"xav\" <= \"lhmq\") and (\"t\" <= \"yqom\") and (\"wkmzt\" > \"dt\") and (\"eru\" > \"r\") and (\"nb\" >= \"vkxyo\") and (\"avsvz\" == \"z\") and (\"weaj\" > \"iazr\") and (\"cdx\" < \"vqrd\") and (\"svlm\" > \"zcvsd\") and (\"n\" == \"sa\") and (\"vt\" > \"qg\") and (\"av\" > \"o\") and (\"hulcd\" != \"k\") and (\"fasnc\" == \"iw\") and (\"t\" < \"vk\") and (\"l\" == \"neh\") and (\"iplv\" != \"fmo\") and (\"g\" < \"lewi\") and (\"xt\" >= \"o\") and (\"uivwz\" == \"fagls\") and (\"oyad\" == \"gri\") and (\"d\" > \"ec\") and (\"rpztf\" >= \"bqtfb\") and (\"hm\" >= \"yfrni\") and (\"pjz\" >= \"gnnd\") and (\"clw\" <= \"zkpd\") and (\"casr\" >= \"pvw\") and (\"vmcvd\" <= \"rvabv\") and (\"zj\" < \"wyfr\") and (\"ki\" <= \"l\")"
	input := a
	program := generateRegisterProgramCombination(input)
	var out interface{}
	var result interface{}
	vm := vm5.New(program)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm.Run(env)
	}
	b.StopTimer()
	result, _ = expr.Eval(input, nil)
	out = vm.Registers[3]
	if out != result {
		b.Fail()
	}
}

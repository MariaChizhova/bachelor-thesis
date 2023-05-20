Provides comparison of different virtual machines for an embedded language in Go:
* tree traversal (evaluator)
* single-stack based virtual machine (vm)
* multiple-stack based virtual machine (vm2, vm4, vm6)
* reflect-based virtual machine (vm3)
* register-based virtual machine (vm5)

### Language definition

#### Values 
Type        | Example                                  |
------------|------------------------------------------|
Integer     | `1` `01`                                 |
Float       | `1.2` `0.1` `.1` `1e2` `1.2e-3` `1.2e+3` |
Bool        | `true` `false`                           |
string      | `"abc"` `'abc'`                          | 
nil         | `nil`                                    | 
array       | `["a", "b", "c"]`                        |

#### Operators:

* Arithmetic: `*`, `/`, `+`, `-`, `%`, `^`
* Comparison: `>`, `<`, `>=`, `<=`, `==`, `!=`
* Logical: `not`, `and`, `or`

#### External:

Can be specified in environment:
* Function call: `map[string]interface{}{"a": 1.2, "b": 2.3}`
* Identifiers: `map[string]interface{}{"a": 1.2, "b": 2.3}`

### How to use it?

```go
	env := map[string]interface{}{"add": func(a, b int64) int64 { return a + b }}
	code := "add(1, 2)"
	tree := parser.Parse(code)
	program, err := compiler.Compile(tree)
	
	vm := vm.New(program.Instructions, program.Constants)
	
	err = vm.Run(env)
	
	if err != nil {
	    panic(err)
	}
	
	out := vm.StackTop()
	fmt.Println(out)
```

In tree traversal you need to call: `out, err = evaluator.Eval(tree, env)`

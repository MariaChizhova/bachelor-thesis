package benchmarks

import (
	"bachelor-thesis/vm5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getSum(n int) string {
	out := "1"
	for i := 2; i <= n; i++ {
		out += "+" + strconv.Itoa(i)
	}
	return out
}

func getExpression(n int) string {
	out := "1"
	for i := 2; i <= n; i++ {
		sign := "+"
		if i%2 == 0 {
			sign = "-"
		}
		out += fmt.Sprintf(" %s %d", sign, i)
	}
	return out
}

func concatenateStrings(num int) string {
	var result strings.Builder
	for i := 0; i < num; i++ {
		result.WriteString(`"` + string('a'+rune(i%26)) + `"`)
		if i != num-1 {
			result.WriteString(" + ")
		}
	}
	return result.String()
}

func concatenateStringsResult(num int) string {
	var result strings.Builder
	for i := 0; i < num; i++ {
		result.WriteString(string('a' + rune(i%26)))
	}
	return result.String()
}

func generateSumBytecode(n int) vm5.Program {
	program := vm5.Program{}
	if n == 1 {
		program.Instructions = []byte{byte(vm5.OpStoreInt), 03, 0}
		program.Constants = []interface{}{1}
	} else {
		program.Instructions = []byte{byte(vm5.OpStoreInt), 01, 0}
		program.Constants = []interface{}{1}
		if n >= 2 {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreInt), 02, 1)
			program.Constants = append(program.Constants, 2)
			program.Instructions = append(program.Instructions, byte(vm5.OpAdd), 03, 01, 02)
		}
		for i := 3; i <= n; i++ {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreInt), 01, byte(i-1))
			program.Constants = append(program.Constants, i)
			program.Instructions = append(program.Instructions, byte(vm5.OpAdd), 03, 01, 03)
		}
	}
	program.Instructions = append(program.Instructions, byte(vm5.OpExit))
	return program
}

func generateExpressionBytecode(n int) vm5.Program {
	program := vm5.Program{}
	if n == 1 {
		program.Instructions = []byte{byte(vm5.OpStoreInt), 03, 0}
		program.Constants = []interface{}{1}
	} else {
		program.Instructions = []byte{byte(vm5.OpStoreInt), 01, 0}
		program.Constants = []interface{}{1}
		if n >= 2 {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreInt), 02, 1)
			program.Constants = append(program.Constants, 2)
			program.Instructions = append(program.Instructions, byte(vm5.OpSub), 03, 01, 02)
		}
		for i := 3; i <= n; i++ {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreInt), 01, byte(i-1))
			program.Constants = append(program.Constants, i)
			if i%2 == 0 {
				program.Instructions = append(program.Instructions, byte(vm5.OpSub), 03, 03, 01)
			} else {
				program.Instructions = append(program.Instructions, byte(vm5.OpAdd), 03, 01, 03)
			}
		}
	}
	program.Instructions = append(program.Instructions, byte(vm5.OpExit))
	return program
}

func generateBytecodeStrings(n int) vm5.Program {
	program := vm5.Program{}
	if n == 1 {
		program.Instructions = []byte{byte(vm5.OpStoreString), 03, 0}
		program.Constants = []interface{}{"a"}
	} else {
		program.Instructions = []byte{byte(vm5.OpStoreString), 01, 0}
		program.Constants = []interface{}{"a"}
		if n >= 2 {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreString), 02, 1)
			program.Constants = append(program.Constants, "b")
			program.Instructions = append(program.Instructions, byte(vm5.OpStringConcat), 03, 01, 02)
		}
		for i := 3; i <= n; i++ {
			program.Instructions = append(program.Instructions, byte(vm5.OpStoreString), 01, byte(i-1))
			program.Constants = append(program.Constants, string('a'+rune((i-1)%26)))
			program.Instructions = append(program.Instructions, byte(vm5.OpStringConcat), 03, 03, 01)
		}
	}
	program.Instructions = append(program.Instructions, byte(vm5.OpExit))
	return program
}

func generateString(numArgs int) string {
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < numArgs; i++ {
		sb.WriteString("(")
		sb.WriteString(generateRandomString())
		sb.WriteString(getRandomComparisonOperator())
		sb.WriteString(generateRandomString())
		sb.WriteString(")")

		if i < numArgs-1 {
			sb.WriteString(getRandomLogicalOperator())
		}
	}

	return sb.String()
}

func generateRandomString() string {
	var sb strings.Builder
	sb.WriteString("\"")
	length := rand.Intn(5) + 1 // length between 1 and 5
	for i := 0; i < length; i++ {
		char := byte(rand.Intn(26) + 'a') // random lowercase letter
		sb.WriteByte(char)
	}
	sb.WriteString("\"")
	return sb.String()
}

func getRandomComparisonOperator() string {
	ops := []string{" > ", " < ", " >= ", " <= ", " == ", " != "}
	return ops[rand.Intn(len(ops))]
}

func getRandomLogicalOperator() string {
	ops := []string{" and ", " or "}
	return ops[rand.Intn(len(ops))]
}

package benchmarks

import (
	"bachelor-thesis/vm5"
	"fmt"
	"math/rand"
	"regexp"
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
	ops := []string{" and "} ///, " or "}
	return ops[rand.Intn(len(ops))]
}

func extractStrings(expr string) []interface{} {
	re := regexp.MustCompile(`"(.+?)"`)
	matches := re.FindAllStringSubmatch(expr, -1)
	strings := make([]interface{}, len(matches))
	for i, match := range matches {
		strings[i] = match[1]
	}
	return strings
}

func generateRegisterProgramCombination(input string) vm5.Program {
	program := vm5.Program{}
	program.Constants = extractStrings(input)
	strIdx := 0
	idx := 1
	parts := strings.Split(input, " ")
	i := 0
	for i+2 < len(parts) {
		program.Instructions = append(program.Instructions, byte(vm5.OpStoreString), 01, byte(strIdx))
		strIdx += 1
		i += 1
		program.Instructions = append(program.Instructions, byte(vm5.OpStoreString), 02, byte(strIdx))
		strIdx += 1
		i += 1
		var opcode int
		switch parts[i-1] {
		case ">":
			opcode = vm5.OpGreaterThan
		case ">=":
			opcode = vm5.OpGreaterOrEqual
		case "<":
			opcode = vm5.OpLessThan
		case "<=":
			opcode = vm5.OpLessOrEqual
		case "==":
			opcode = vm5.OpEqual
		case "!=":
			opcode = vm5.OpNotEqual
		}
		program.Instructions = append(program.Instructions, byte(opcode), 03, 01, 02)
		i += 1
		if i >= len(parts) {
			break
		}
		idx += 1
		if parts[i] == "or" {
			program.Instructions = append(program.Instructions, byte(vm5.OpJumpIfTrue), 03, byte(13*idx-3))
			i += 1
		} else if parts[i] == "and" {
			program.Instructions = append(program.Instructions, byte(vm5.OpJumpIfFalse), 03, byte(13*idx-3))
			i += 1
		}
	}
	program.Instructions = append(program.Instructions, byte(vm5.OpExit))
	return program
}

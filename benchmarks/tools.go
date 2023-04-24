package benchmarks

import (
	"bachelor-thesis/vm4"
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

func generateSumBytecode(n int) []int64 {
	bytecode := []int64{
		vm4.OpConstant, vm4.R0, 1,
	}
	for i := 2; i <= n; i++ {
		bytecode = append(bytecode, vm4.OpConstant, vm4.R1, int64(i))
		bytecode = append(bytecode, vm4.OpAdd, vm4.R0, vm4.R1)
	}
	bytecode = append(bytecode, vm4.OpPrint, vm4.R0, vm4.OpHalt)
	return bytecode
}

func generateSumBytecodeInterface(n int) []interface{} {
	bytecode := []interface{}{
		vm4.OpConstant, vm4.R0, int64(1),
	}
	for i := 2; i <= n; i++ {
		bytecode = append(bytecode, vm4.OpConstant, vm4.R1, int64(i))
		bytecode = append(bytecode, vm4.OpAdd, vm4.R0, vm4.R1)
	}
	bytecode = append(bytecode, vm4.OpPrint, vm4.R0, vm4.OpHalt)
	return bytecode
}

func generateSumBytecodeInterfaceStrings(n int) []interface{} {
	bytecode := []interface{}{
		vm4.OpConstant, vm4.R0, "a",
	}
	for i := 1; i < n; i++ {
		bytecode = append(bytecode, vm4.OpConstant, vm4.R1, string('a'+rune(i%26)))
		bytecode = append(bytecode, vm4.OpAdd, vm4.R0, vm4.R1)
	}
	bytecode = append(bytecode, vm4.OpPrint, vm4.R0, vm4.OpHalt)
	return bytecode
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

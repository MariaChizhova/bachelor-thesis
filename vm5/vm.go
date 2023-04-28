package vm5

import (
	"fmt"
	"math"
	"reflect"
)

type VM struct {
	Registers    [16]interface{}
	ip           int
	instructions []byte
	constants    []interface{}
}

func New(program Program) *VM {
	return &VM{
		instructions: program.Instructions,
		constants:    program.Constants,
	}
}

func (vm *VM) Run(env interface{}) error {
	vm.ip = 0
	run := true
	for run {
		op := NewOpcode(vm.instructions[vm.ip])
		switch int(op.Value()) {
		case OpExit:
			run = false
		case OpStoreInt:
			vm.ip++
			reg := int(vm.instructions[vm.ip])
			if reg >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", reg)
			}
			vm.ip++
			val := vm.constants[vm.instructions[vm.ip]]
			vm.ip++
			vm.Registers[reg] = val
		case OpAdd:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++
			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal + bVal
		case OpSub:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal - bVal
		case OpMul:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal * bVal
		case OpDiv:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal / bVal
		case OpMod:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = aVal % bVal
		case OpExp:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			vm.ip++
			b := vm.instructions[vm.ip]
			vm.ip++

			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(int)
			bVal := vm.Registers[b].(int)
			vm.Registers[res] = int(math.Pow(float64(aVal), float64(bVal)))
		case OpStoreString:
			vm.ip++
			reg := vm.instructions[vm.ip]
			if int(reg) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", reg)
			}
			vm.ip++
			str := vm.constants[vm.instructions[vm.ip]]
			vm.ip++
			vm.Registers[reg] = str
		case OpStringConcat:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			a := vm.instructions[vm.ip]
			if int(a) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", a)
			}
			vm.ip++
			b := vm.instructions[vm.ip]
			if int(b) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", b)
			}
			vm.ip++
			if int(res) >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", res)
			}
			aVal := vm.Registers[a].(string)
			bVal := vm.Registers[b].(string)
			vm.Registers[res] = aVal + bVal
		case OpStoreBool:
			vm.ip++
			reg := int(vm.instructions[vm.ip])
			if reg >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", reg)
			}
			vm.ip++
			val := int(vm.instructions[vm.ip])
			vm.ip++
			vm.Registers[reg] = val != 0
		case OpEqual:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal == bVal
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal == bVal
			case bool:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal == bVal
			}
		case OpNotEqual:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal != bVal
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal != bVal
			case bool:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal != bVal
			}
		case OpLessThan:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(int) < bVal.(int)
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(string) < bVal.(string)
			}
		case OpGreaterThan:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(int) > bVal.(int)
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(string) > bVal.(string)
			}
		case OpLessOrEqual:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(int) <= bVal.(int)
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(string) <= bVal.(string)
			}
		case OpGreaterOrEqual:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			r1 := int(vm.instructions[vm.ip])
			vm.ip++
			r2 := int(vm.instructions[vm.ip])
			vm.ip++

			if r1 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r1)
			}
			if r2 >= len(vm.Registers) {
				return fmt.Errorf("register %d out of range", r2)
			}

			switch vm.Registers[r1].(type) {
			case int:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(int) >= bVal.(int)
			case string:
				aVal := vm.Registers[r1]
				bVal := vm.Registers[r2]
				vm.Registers[res] = aVal.(string) >= bVal.(string)
			}
		case OpCall:
			vm.ip++
			res := vm.instructions[vm.ip]
			vm.ip++
			fnAddr := vm.constants[vm.instructions[vm.ip]]
			vm.ip++
			v := reflect.ValueOf(env)
			kind := v.Kind()
			if kind == reflect.Invalid {
				fmt.Sprintf("error")
			}
			var fn interface{}
			switch kind {
			case reflect.Map:
				value := v.MapIndex(reflect.ValueOf(fnAddr))
				if value.IsValid() {
					fn = value.Interface()
				} else {
					elem := reflect.TypeOf(env).Elem()
					fn = reflect.Zero(elem).Interface()
				}
			}
			size := int(vm.instructions[vm.ip])
			in := make([]reflect.Value, size)
			vm.ip++
			for i := 0; i < size; i++ {
				r := int(vm.instructions[vm.ip])
				vm.ip++
				in[i] = reflect.ValueOf(vm.Registers[r])
			}
			out := reflect.ValueOf(fn).Call(in)
			vm.Registers[res] = out[0].Interface()
		}
	}
	return nil
}

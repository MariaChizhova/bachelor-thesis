package evaluator

import (
	"bachelor-thesis/parser/ast"
	"fmt"
	"math"
	"reflect"
)

func Eval(node ast.Node, env interface{}) (interface{}, error) {
	switch node.Type() {
	case ast.NodeNumber:
		return EvalNumber(node)
	case ast.NodeIdentifier:
		v := reflect.ValueOf(env)
		return v.MapIndex(reflect.ValueOf(node.(*ast.IdentifierNode).Value)).Interface(), nil
	case ast.NodeString:
		return node.(*ast.StringNode).Value, nil
	case ast.NodeBool:
		return node.(*ast.BoolNode).Value, nil
	case ast.NodeNil:
		return nil, nil
	case ast.NodeUnary:
		return EvalUnary(node, env)
	case ast.NodeBinary:
		return EvalBinary(node, env)
	case ast.NodeCall:
		return EvalFunctionCall(node, env)
	case ast.NodeArray:
		return EvalArray(node, env)
	case ast.NodeMember:
		return EvalIndex(node, env)
	}
	return nil, nil
}

func EvalNumber(node ast.Node) (interface{}, error) {
	if node.(*ast.NumberNode).IsInt {
		return node.(*ast.NumberNode).Int64, nil
	} else if node.(*ast.NumberNode).IsFloat {
		return node.(*ast.NumberNode).Float64, nil
	}
	return nil, nil
}

func EvalUnary(node ast.Node, env interface{}) (interface{}, error) {
	value, err := Eval(node.(*ast.UnaryNode).Node, env)
	if err != nil {
		return nil, err
	}
	switch node.(*ast.UnaryNode).Operator {
	case "+":
		switch t := value.(type) {
		case int64:
			return t, nil
		case float64:
			return t, nil
		default:
			fmt.Printf("unexpected type %T", t)
		}
	case "-":
		switch t := value.(type) {
		case int64:
			return -t, nil
		case float64:
			return -t, nil
		default:
			fmt.Printf("unexpected type %T", t)
		}
	case "not":
		return !value.(bool), nil
	}
	return nil, fmt.Errorf("undefined unary %q operator", node.(*ast.UnaryNode).Operator)
}

func EvalBinary(node ast.Node, env interface{}) (interface{}, error) {
	left, err := Eval(node.(*ast.BinaryNode).Left, env)
	if err != nil {
		return nil, err
	}
	right, err := Eval(node.(*ast.BinaryNode).Right, env)
	if err != nil {
		return nil, err
	}
	switch node.(*ast.BinaryNode).Operator {
	case "or":
		switch l := left.(type) {
		case bool:
			switch r := right.(type) {
			case bool:
				return l || r, nil
			default:
				return nil, fmt.Errorf("non-bool value in cond (%T)", r)
			}
		default:
			return nil, fmt.Errorf("non-bool value in cond (%T)", l)
		}
	case "and":
		switch l := left.(type) {
		case bool:
			switch r := right.(type) {
			case bool:
				return l && r, nil
			default:
				return nil, fmt.Errorf("non-bool value in cond (%T)", r)
			}
		default:
			return nil, fmt.Errorf("non-bool value in cond (%T)", l)
		}
	case "+":
		switch l := left.(type) {
		case int:
			switch r := right.(type) {
			case int:
				return l + r, nil
			}
		case int64:
			switch r := right.(type) {
			case int64:
				return l + r, nil
			case float64:
				return float64(l) + r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l + float64(r), nil
			case float64:
				return l + r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l + r, nil
			}
		}
	case "-":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l - r, nil
			case float64:
				return float64(l) - r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l - float64(r), nil
			case float64:
				return l - r, nil
			}
		}
	case "*":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l * r, nil
			case float64:
				return float64(l) * r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l * float64(r), nil
			case float64:
				return l * r, nil
			}
		}
	case "/":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l / r, nil
			case float64:
				return float64(l) / r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l / float64(r), nil
			case float64:
				return l / r, nil
			}
		}
	case "%":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l % r, nil
			}
		}
	case "^":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return int64(math.Pow(float64(l), float64(r))), nil
			case float64:
				return math.Pow(float64(l), r), nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return math.Pow(l, float64(r)), nil
			case float64:
				return math.Pow(l, r), nil
			}
		}
	case "<":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l < r, nil
			case float64:
				return float64(l) < r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l < float64(r), nil
			case float64:
				return l < r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l < r, nil
			}
		}
	case "<=":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l <= r, nil
			case float64:
				return float64(l) <= r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l <= float64(r), nil
			case float64:
				return l <= r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l <= r, nil
			}
		}
	case ">":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l > r, nil
			case float64:
				return float64(l) > r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l > float64(r), nil
			case float64:
				return l > r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l > r, nil
			}
		}
	case ">=":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return l >= r, nil
			case float64:
				return float64(l) >= r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l >= float64(r), nil
			case float64:
				return l >= r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l >= r, nil
			}
		}
	case "==":
		switch l := left.(type) {
		case bool:
			switch r := right.(type) {
			case bool:
				return l == r, nil
			}
		case int64:
			switch r := right.(type) {
			case int64:
				return l == r, nil
			case float64:
				return float64(l) == r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l == float64(r), nil
			case float64:
				return l == r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l == r, nil
			}
		}
	case "!=":
		switch l := left.(type) {
		case bool:
			switch r := right.(type) {
			case bool:
				return l != r, nil
			}
		case int64:
			switch r := right.(type) {
			case int64:
				return l != r, nil
			case float64:
				return float64(l) != r, nil
			}
		case float64:
			switch r := right.(type) {
			case int64:
				return l != float64(r), nil
			case float64:
				return l != r, nil
			}
		case string:
			switch r := right.(type) {
			case string:
				return l != r, nil
			}
		}
	}
	return nil, fmt.Errorf("undefined binary %q operator", node.(*ast.BinaryNode).Operator)
}

func EvalArray(node ast.Node, env interface{}) (interface{}, error) {
	array := make([]interface{}, 0)
	for _, node := range node.(*ast.ArrayNode).Nodes {
		value, err := Eval(node, env)
		if err != nil {
			return nil, err
		}
		array = append(array, value)
	}
	return array, nil
}

func EvalIndex(node ast.Node, env interface{}) (interface{}, error) {
	array, err := Eval(node.(*ast.MemberNode).Node, env)
	if err != nil {
		return array, err
	}
	index, err := Eval(node.(*ast.MemberNode).Property, env)
	if err != nil {
		return index, err
	}
	arrayObject := array.([]interface{})
	i := index.(int64)
	max := int64(len(arrayObject) - 1)
	if i < 0 || i > max {
		return nil, nil
	}
	switch t := arrayObject[i].(type) {
	case int64, float64, bool, string:
		return t, nil
	}
	return nil, nil
}

func EvalFunctionCall(node ast.Node, env interface{}) (interface{}, error) {
	node = node.(*ast.CallNode)
	name := node.(*ast.CallNode).Callee.(*ast.IdentifierNode).Value
	fn, ok := getFunc(env, name)
	if !ok {
		return nil, fmt.Errorf("undefined: %v", name)
	}

	in := make([]reflect.Value, 0)

	for _, a := range node.(*ast.CallNode).Arguments {
		i, err := Eval(a, env)
		if err != nil {
			return nil, err
		}
		in = append(in, reflect.ValueOf(i))
	}

	out := reflect.ValueOf(fn).Call(in)

	if len(out) == 0 {
		return nil, nil
	} else if len(out) > 1 {
		return nil, fmt.Errorf("func %q must return only one value", name)
	}

	if out[0].IsValid() && out[0].CanInterface() {
		return out[0].Interface(), nil
	}
	return nil, nil
}

func getFunc(val interface{}, i interface{}) (interface{}, bool) {
	v := reflect.ValueOf(val)
	d := v
	if v.Kind() == reflect.Ptr {
		d = v.Elem()
	}

	switch d.Kind() {
	case reflect.Map:
		value := d.MapIndex(reflect.ValueOf(i))
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	case reflect.Struct:
		name := reflect.ValueOf(i).String()
		method := v.MethodByName(name)
		if method.IsValid() && method.CanInterface() {
			return method.Interface(), true
		}
		value := d.FieldByName(name)
		if value.IsValid() && value.CanInterface() {
			return value.Interface(), true
		}
	}
	return nil, false
}

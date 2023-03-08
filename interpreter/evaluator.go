package interpreter

import (
	"bachelor-thesis/parser/ast"
	"fmt"
	"math"
)

func Eval(node ast.Node) (interface{}, error) {
	switch node.Type() {
	case ast.NodeNumber:
		return EvalNumber(node)
	case ast.NodeIdentifier:
		return node.(*ast.IdentifierNode).Value, nil
	case ast.NodeString:
		return node.(*ast.StringNode).Value, nil
	case ast.NodeBool:
		return node.(*ast.BoolNode).Value, nil
	case ast.NodeNil:
		return nil, nil
	case ast.NodeUnary:
		return EvalUnary(node)
	case ast.NodeBinary:
		return EvalBinary(node)
	case ast.NodeFunction:
		//return
	case ast.NodeArray:
		return EvalArray(node)
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

func EvalUnary(node ast.Node) (interface{}, error) {
	value, err := Eval(node.(*ast.UnaryNode).Node)
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

func EvalBinary(node ast.Node) (interface{}, error) {
	left, err := Eval(node.(*ast.BinaryNode).Left)
	if err != nil {
		return nil, err
	}
	right, err := Eval(node.(*ast.BinaryNode).Right)
	if err != nil {
		return nil, err
	}
	switch node.(*ast.BinaryNode).Operator {
	case "+":
		switch l := left.(type) {
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
		case int:
			switch r := right.(type) {
			case int64:
				return int64(l) % r, nil
			}
		}
	case "^":
		switch l := left.(type) {
		case int64:
			switch r := right.(type) {
			case int64:
				return int(math.Pow(float64(l), float64(r))), nil
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

func EvalArray(node ast.Node) (interface{}, error) {
	array := make([]interface{}, 0)
	for _, node := range node.(*ast.ArrayNode).Nodes {
		value, err := Eval(node)
		if err != nil {
			return nil, err
		}
		array = append(array, value)
	}
	return array, nil
}

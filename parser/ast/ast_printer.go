package ast

import (
	"fmt"
	"reflect"
	"unicode"
	"unicode/utf8"
)

func printAST(v reflect.Value, indent string) string {
	if !v.IsValid() {
		return "nil"
	}
	t := v.Type()
	switch t.Kind() {
	case reflect.Ptr:
		return printAST(v.Elem(), indent)
	case reflect.Interface:
		return printAST(v.Elem(), indent)
	case reflect.String:
		return fmt.Sprintf("%q", v)
	case reflect.Slice:
		out := t.String() + "{"
		var body string
		if v.Len() > 0 {
			body += "\n"
			for i := 0; i < v.Len(); i++ {
				body += fmt.Sprintf("%v%v,", indent+"\t", printAST(v.Index(i), indent+"\t"))
				if i+1 < v.Len() {
					body += "\n"
				}
			}
			body += "\n" + indent
		}
		return out + body + "}"
	case reflect.Struct:
		out := t.Name() + " {\n"
		for i := 0; i < t.NumField(); i++ {
			if name := t.Field(i).Name; IsExported(name) {
				value := v.Field(i)
				out += fmt.Sprintf("%v%v: %v,\n", indent+"\t", name, printAST(value, indent+"\t"))
			}
		}
		return out + indent + "}"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func IsExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

func Print(node Node) string {
	return printAST(reflect.ValueOf(node), "")
}

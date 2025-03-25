package ast

import (
	"fmt"
	"go/ast"
	"reflect"
)

type printVisitor struct {
}

func (p *printVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		fmt.Println(nil)
		return p
	}

	typ := reflect.TypeOf(node)
	val := reflect.ValueOf(node)

	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	fmt.Printf("type: %s, val: %+v\n", typ.Name(), val.Interface())
	return p
}

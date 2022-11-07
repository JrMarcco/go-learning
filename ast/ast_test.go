package ast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var src = `
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
`

func TestAst(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	ast.Walk(&printVisitor{}, f)
}

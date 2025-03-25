package ast

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

var fileCase = `
// annotation go through the source code and extra the annotation
// @author jrmarcco
/* @multiple first line
second line
*/
// @date 2022/10/22
package annotation
`

var funcTypeCase = `
// annotation go through the source code and extra the annotation
// @author jrmarcco
/* @multiple first line
second line
*/
// @date 2022/10/22
package annotation

type (
	// FuncType is a type
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	FuncType func()
)
`

var structTypeCase = `
// annotation go through the source code and extra the annotation
// @author jrmarcco
/* @multiple first line
second line
*/
// @date 2022/10/22
package annotation

type (
	// StructType is a test struct
	//
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	StructType struct {
		// Public is a field
		// @type string
		Public string
	}

	// SecondType is a test struct
	//
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	SecondType struct {
	}
)
`

var multiTypeCase = `
// annotation go through the source code and extra the annotation
// @author jrmarcco
/* @multiple first line
second line
*/
// @date 2022/10/22
package annotation

type (
	// FuncType is a type
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	FuncType func()
)

type (
	// StructType is a test struct
	//
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	StructType struct {
		// Public is a field
		// @type string
		Public string
	}

	// SecondType is a test struct
	//
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	SecondType struct {
	}
)

type (
	// Interface is a test interface
	// @author jrmarcco
	/* @multiple first line
	   second line
	*/
	// @date 2022/10/22
	Interface interface {
		// MyFunc is a test func
		// @parameter arg1 int
		// @parameter arg2 int32
		// @return string
		MyFunc(arg1 int, arg2 int32) string

		// second is a test func
		// @return string
		second() string
	}
)
`

func TestAstFile(t *testing.T) {
	tcs := []struct {
		name    string
		arg     string
		wantRes File
	}{
		{
			name: "package annotations",
			arg:  fileCase,
			wantRes: File{
				Annos: []Annotation{
					{"author", "jrmarcco"},
					{"multiple", "first line\nsecond line\n"},
					{"date", "2022/10/22"},
				},
				Types: []Type{},
			},
		},
		{
			name: "func type annotations",
			arg:  funcTypeCase,
			wantRes: File{
				Annos: []Annotation{
					{"author", "jrmarcco"},
					{"multiple", "first line\nsecond line\n"},
					{"date", "2022/10/22"},
				},
				Types: []Type{
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
					},
				},
			},
		},
		{
			name: "struct type annotations",
			arg:  structTypeCase,
			wantRes: File{
				Annos: []Annotation{
					{"author", "jrmarcco"},
					{"multiple", "first line\nsecond line\n"},
					{"date", "2022/10/22"},
				},
				Types: []Type{
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
						Fields: []Field{
							{
								Annos: []Annotation{
									{"type", "string"},
								},
							},
						},
					},
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
					},
				},
			},
		},
		{
			name: "multi type annotations",
			arg:  multiTypeCase,
			wantRes: File{
				Annos: []Annotation{
					{"author", "jrmarcco"},
					{"multiple", "first line\nsecond line\n"},
					{"date", "2022/10/22"},
				},
				Types: []Type{
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
					},
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
						Fields: []Field{
							{
								Annos: []Annotation{
									{"type", "string"},
								},
							},
						},
					},
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
					},
					{
						Annos: []Annotation{
							{"author", "jrmarcco"},
							{"multiple", "first line\n\t   second line\n\t"},
							{"date", "2022/10/22"},
						},
						Fields: []Field{
							{
								Annos: []Annotation{
									{"parameter", "arg1 int"},
									{"parameter", "arg2 int32"},
									{"return", "string"},
								},
							},
							{
								Annos: []Annotation{
									{"return", "string"},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", tc.arg, parser.ParseComments)
			if err != nil {
				t.Fatal(err)
			}

			ve := &VisitorEntry{}
			ast.Walk(ve, f)

			assert.Equal(t, tc.wantRes, ve.Get())
		})
	}
}

package ast

import (
	"go/ast"
)

type VisitorEntry struct {
	*fileVisitor
}

func (v *VisitorEntry) Visit(node ast.Node) ast.Visitor {
	file, ok := node.(*ast.File)
	if ok {
		v.fileVisitor = &fileVisitor{
			annotations: buildAnnos(file.Doc),
		}
		return v.fileVisitor
	}
	return v
}

func (v *VisitorEntry) Get() File {
	return v.get()
}

type fileVisitor struct {
	annotations annotations
	types       []*typeSpecVisitor
}

func (f *fileVisitor) Visit(node ast.Node) ast.Visitor {
	typeSpec, ok := node.(*ast.TypeSpec)
	if ok {
		tsv := &typeSpecVisitor{
			annotations: buildAnnos(typeSpec.Doc),
		}
		f.types = append(f.types, tsv)
		return tsv
	}
	return f
}

type File struct {
	Annos []Annotation
	Types []Type
}

func (f *fileVisitor) get() File {
	types := make([]Type, 0, len(f.types))
	for _, typ := range f.types {
		types = append(types, typ.get())
	}

	return File{
		Annos: f.annotations.annos,
		Types: types,
	}
}

type typeSpecVisitor struct {
	annotations annotations
	fields      []Field
}

func (t *typeSpecVisitor) Visit(node ast.Node) ast.Visitor {
	field, ok := node.(*ast.Field)
	if ok {
		annotations := buildAnnos(field.Doc)
		if len(annotations.annos) != 0 {
			t.fields = append(t.fields, Field{Annos: annotations.annos})
		}
	}
	return t
}

type Type struct {
	Annos  []Annotation
	Fields []Field
}

type Field struct {
	Annos []Annotation
}

func (t *typeSpecVisitor) get() Type {
	return Type{
		Annos:  t.annotations.annos,
		Fields: t.fields,
	}
}

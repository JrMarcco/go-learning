package ast

import (
	"go/ast"
	"strings"
)

type annotations struct {
	annos []Annotation
}

type Annotation struct {
	Key string
	Val string
}

func buildAnnos(cg *ast.CommentGroup) annotations {
	if cg == nil || len(cg.List) == 0 {
		return annotations{}
	}

	annos := make([]Annotation, 0, len(cg.List))
	for _, doc := range cg.List {
		text, ok := extractAnnotation(doc.Text)
		if !ok {
			continue
		}
		if strings.HasPrefix(text, "@") {
			segs := strings.SplitN(text[1:], " ", 2)
			if len(segs) < 2 {
				continue
			}

			annos = append(annos, Annotation{
				Key: segs[0],
				Val: segs[1],
			})
		}
	}

	return annotations{
		annos: annos,
	}
}

func extractAnnotation(text string) (string, bool) {
	if strings.HasPrefix(text, "//") {
		return strings.TrimLeft(text[2:], " "), true
	}
	if strings.HasPrefix(text, "/*") {
		return strings.Trim(text[2:len(text)-2], " "), true
	}
	return "", false
}

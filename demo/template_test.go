package demo

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"text/template"
)

func TestBasic(t *testing.T) {

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(`Hello, {{ . }}`)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, "jrmarcco")
	require.NoError(t, err)

	assert.Equal(t, "Hello, jrmarcco", bb.String())
}

func TestStruct(t *testing.T) {

	user := struct {
		Name string
	}{
		Name: "jrmarcco",
	}

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(`Hello, {{ .Name }}`)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	// 这里 user 传入结构体还是指针没有影响
	err = tpl.Execute(bb, &user)
	require.NoError(t, err)

	assert.Equal(t, "Hello, jrmarcco", bb.String())

}

func TestMapData(t *testing.T) {

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(`Hello, {{ .Name }}`)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, map[string]string{"Name": "jrmarcco"})
	require.NoError(t, err)

	assert.Equal(t, "Hello, jrmarcco", bb.String())
}

func TestSlice(t *testing.T) {

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(`Hello, {{ index . 0 }}`)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, []string{"jrmarcco"})
	require.NoError(t, err)

	assert.Equal(t, "Hello, jrmarcco", bb.String())
}

func TestVariableDeclare(t *testing.T) {

	// '-': 去除空格和换行符
	testTpl := `
		{{- $userName := index . 0 -}}
		Hello, {{ $userName -}}
	`

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(testTpl)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, []string{"jrmarcco"})
	require.NoError(t, err)

	assert.Equal(t, "Hello, jrmarcco", bb.String())
}

func TestFuncCall(t *testing.T) {

	testTpl := `
fields lens: {{ len .Fields }}
{{ .Hello .Name}}
{{ printf "%.2f" 1.009}}
`

	tpl := template.New("hello-world")

	tpl, err := tpl.Parse(testTpl)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, &Func{
		Name:   "jrmarcco",
		Fields: []string{"name", "age", "email"},
	})
	require.NoError(t, err)

	assert.Equal(t,
		`
fields lens: 3
Hello, jrmarcco
1.01
`,
		bb.String(),
	)
}

type Func struct {
	Name   string
	Fields []string
}

func (f *Func) Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func TestRange(t *testing.T) {
	testTpl := `
{{- range $index, $elem := .Fields -}}
{{- . }}
{{ $index }} - {{ $elem }}
{{ end -}}
`
	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(testTpl)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, Func{Fields: []string{"Zero", "One", "Two", "Three"}})
	require.NoError(t, err)

	assert.Equal(t, `Zero
0 - Zero
One
1 - One
Two
2 - Two
Three
3 - Three
`, bb.String())
}

func TestFor(t *testing.T) {
	testTpl := `
{{- range $index, $ignore := . }}
{{- $index -}}
{{ end -}}
`

	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(testTpl)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, make([]int, 10))
	require.NoError(t, err)

	assert.Equal(t, "0123456789", bb.String())
}

func TestIfElse(t *testing.T) {
	testTpl := `
{{- range $index, $elem := . -}}
{{ if and (gt $elem 0) (le $elem 2) -}}
small
{{ else if and (gt $elem 2) (le $elem 4) -}}
medium
{{ else if and (gt $elem 4) (le $elem 6) -}}
large
{{ end -}}
{{ end -}}
`

	tpl := template.New("hello-world")
	tpl, err := tpl.Parse(testTpl)
	require.NoError(t, err)

	bb := &bytes.Buffer{}
	err = tpl.Execute(bb, []int{1, 2, 3, 4, 5, 6})
	require.NoError(t, err)

	assert.Equal(t, `small
small
medium
medium
large
large
`, bb.String())
}

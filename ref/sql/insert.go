package sql

import (
	"errors"
	"reflect"
	"strings"
	"unicode"
)

var (
	NilEntityErr     = errors.New("nil entity")
	InvalidEntityErr = errors.New("invalid entity")
)

func InsertStmt(entity any) (string, []any, error) {

	if entity == nil {
		return "", nil, NilEntityErr
	}

	val := reflect.ValueOf(entity)
	typ := val.Type()

	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return "", nil, InvalidEntityErr
	}

	bd := strings.Builder{}
	bd.WriteString("INSERT INTO `")
	bd.WriteString(camelToUnderLine(typ.Name()))
	bd.WriteString("`(")

	fdNames, fdVals := fdNamesAndVals(val)
	args := make([]any, 0, len(fdVals))
	suffixBd := strings.Builder{}

	for index, fdName := range fdNames {
		if index > 0 {
			bd.WriteRune(',')
			suffixBd.WriteRune(',')
		}
		bd.WriteRune('`')
		bd.WriteString(fdName)
		bd.WriteRune('`')

		args = append(args, fdVals[fdName])
		suffixBd.WriteRune('?')
	}
	bd.WriteString(") VALUES (")
	bd.WriteString(suffixBd.String())
	bd.WriteString(");")

	return bd.String(), args, nil
}

func fdNamesAndVals(val reflect.Value) ([]string, map[string]any) {
	typ := val.Type()
	fdCnt := typ.NumField()

	fdNames := make([]string, 0, fdCnt)
	fdVals := make(map[string]any, fdCnt)

	for i := 0; i < fdCnt; i++ {
		fd := typ.Field(i)
		fdVal := val.Field(i)

		// 处理组合
		if fd.Type.Kind() == reflect.Struct && fd.Anonymous {
			sFdNames, sFdVals := fdNamesAndVals(fdVal)

			for _, sFdName := range sFdNames {
				if _, ok := fdVals[sFdName]; ok {
					continue
				}
				fdNames = append(fdNames, sFdName)
				fdVals[sFdName] = sFdVals[sFdName]
			}
			continue
		}

		fdName := camelToUnderLine(fd.Name)

		fdNames = append(fdNames, fdName)
		fdVals[fdName] = fdVal.Interface()
	}
	return fdNames, fdVals
}

func camelToUnderLine(src string) string {
	if src == "" {
		return src
	}

	bd := strings.Builder{}
	for i, r := range src {
		if unicode.IsUpper(r) {
			if i > 0 {
				bd.WriteRune('_')
			}
			bd.WriteRune(unicode.ToLower(r))
		} else {
			bd.WriteRune(r)
		}
	}
	return bd.String()
}

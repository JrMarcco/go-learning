package ref

import (
	"fmt"
	"log"
	"reflect"
)

func IterateFields(entity any) {
	fm, err := iterateFields(entity)
	if err != nil {
		log.Fatalln(err)
	}

	for name, fd := range fm {
		fmt.Printf("%s: %s \n", name, fd)
	}
}

func iterateFields(entity any) (map[string]any, error) {
	if entity == nil {
		return nil, NilErr
	}

	typ := reflect.TypeOf(entity)
	val := reflect.ValueOf(entity)

	if val.IsZero() {
		return nil, ZeroValErr
	}

	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Pointer {
		return nil, InvalidTypErr
	}

	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
		val = val.Elem()
	}

	// NumField 在 struct 以外的类型上调用会直接 panic
	fdCnt := typ.NumField()
	fm := make(map[string]any, fdCnt)

	for i := 0; i < fdCnt; i++ {
		fd := typ.Field(i)
		if fd.IsExported() {
			fdk := fd.Type.Kind()
			if fdk == reflect.Struct || fdk == reflect.Pointer {
				fdVal := val.Field(i)
				if !fdVal.IsZero() {
					sfm, _ := iterateFields(val.Field(i).Interface())
					fm[fd.Name] = sfm
				}
			} else {
				fdVal := val.Field(i)
				fm[fd.Name] = fdVal.Interface()
			}
		}
	}
	return fm, nil
}

func SetField(entity any, field string, newVal any) error {

	if entity == nil {
		return NilErr
	}

	typ := reflect.TypeOf(entity)
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return InvalidTypErr
	}

	typ = typ.Elem()
	val := reflect.ValueOf(entity).Elem()

	if _, exist := typ.FieldByName(field); !exist {
		return NotExistFieldErr
	}

	fd := val.FieldByName(field)
	if !fd.CanSet() {
		return UnsetFiledErr
	}

	fd.Set(reflect.ValueOf(newVal))

	return nil
}

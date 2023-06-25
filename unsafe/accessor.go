package unsafe

import (
	"errors"
	"reflect"
	"unsafe"
)

/*
Read:
	*(*T)(ptr)
	T 是目标类型
Write:
	*(*T)(ptr) = T
	T 是目标类型

当目标类型未知时可以使用 reflect.NewAt(typ, ptr).Elem()。
*/

type FieldAccessor struct {
	fm   map[string]fieldMeta // 字段元数据
	addr unsafe.Pointer       // 对象起始地址
}

type fieldMeta struct {
	offset uintptr
	typ    reflect.Type
}

// NewAccessor
// t 只能是一级指针
func NewAccessor(t any) (*FieldAccessor, error) {
	if t == nil {
		return nil, errors.New("can not be nil")
	}

	typ := reflect.TypeOf(t)
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("invalid type")
	}

	typ = typ.Elem()
	numField := typ.NumField()
	fm := make(map[string]fieldMeta, numField)

	for i := 0; i < numField; i++ {
		fd := typ.Field(i)
		fm[fd.Name] = fieldMeta{
			offset: fd.Offset,
			typ:    fd.Type,
		}
	}

	return &FieldAccessor{
		fm:   fm,
		addr: reflect.ValueOf(t).UnsafePointer(),
	}, nil
}

// Get 访问具体字段获取字段值
//
// *(*T)(ptr) : T 是目标类型
// 目标类型未知时使用 reflect.NewAt(typ, ptr).Elem()。
func (f *FieldAccessor) Get(fdName string) (any, error) {
	fm, ok := f.fm[fdName]
	if !ok {
		return nil, errors.New("invalid field")
	}

	ptr := unsafe.Pointer(uintptr(f.addr) + fm.offset)
	if ptr == nil {
		return nil, errors.New("invalid field addr")
	}

	return reflect.NewAt(fm.typ, ptr).Elem().Interface(), nil
}

// Set 访问具体字段并修改其字段值
//
// *(*T)(ptr) : T 是目标类型
// 当目标类型未知时可以使用 reflect.NewAt(typ, ptr).Elem()。
func (f *FieldAccessor) Set(fdName string, val any) error {
	fm, ok := f.fm[fdName]
	if !ok {
		return errors.New("invalid field")
	}

	ptr := unsafe.Pointer(uintptr(f.addr) + fm.offset)
	if ptr == nil {
		return errors.New("invalid field addr")
	}

	reflect.NewAt(fm.typ, ptr).Elem().Set(reflect.ValueOf(val))
	return nil
}
